<script lang="ts">
	import 'carbon-components-svelte/css/all.css';
	import Notification from '../components/Notification.svelte';
	import LL from '$lib/i18n/i18n-svelte';
	import {
		Header,
		HeaderUtilities,
		HeaderAction,
		HeaderGlobalAction,
		HeaderPanelLinks,
		HeaderPanelDivider,
		HeaderPanelLink,
		SideNav,
		SideNavItems,
		SideNavMenuItem,
		SideNavLink,
		SkipToContent,
		Content,
		Grid,
		Row,
		Column,
		HeaderActionLink,
		SideNavMenu,
		ProgressBar,
		Theme
	} from 'carbon-components-svelte';
	import SettingsAdjust from 'carbon-icons-svelte/lib/SettingsAdjust.svelte';
	import UserAvatarFilledAlt from 'carbon-icons-svelte/lib/UserAvatarFilledAlt.svelte';
	import LogoGithub from 'carbon-icons-svelte/lib/LogoGithub.svelte';
	import session from '$lib/stores/userStore';
	import setup from '$lib/stores/setupStore';
	import { page } from '$app/stores';
	import { SessionRole } from '$lib/models/Session';
	import { goto } from '$app/navigation';
	import { sendSuccess } from '$lib/util/notifications';
	import { onMount, tick } from 'svelte';
	import { handleResponse } from '$lib/util/handleResponse';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { API_URL, SET_UP } from '$lib/api/constants';

	let isSideNavOpen = false;
	let isUserMenuOpen = false;
	let canRender = true;

	//Using the stores should guarantee that the request
	//is sent only once

	onMount(function () {
		if (!$setup) {
			canRender = false;
		}
	});

	async function checkSetup() {
		const response = await fetchWithRefresh(API_URL + SET_UP, { method: 'GET' });

		if (response.ok) {
			$setup = true;
		}

		handleResponse(response.status, false);
		await tick();

		canRender = true;
	}

	async function handleSignOut() {
		//TODO check if this try catch is needed
		try {
			await fetch('/api/sign-out', { method: 'POST', credentials: 'include' });

			//Set session data to false and redirect to the home page
			$session.active = false;
			$session.role = SessionRole.None;

			sendSuccess('Signed out correctly', 'Thank you for using Libre-ASI');

			goto('/');
		} catch (e) {
			console.error(e);
		} finally {
			isUserMenuOpen = false;
		}
	}
</script>

<Theme persist persistKey="__carbon-theme" />

<Header bind:isSideNavOpen platformName="Libre-ASI" href="/">
	<svelte:fragment slot="skip-to-content">
		<SkipToContent />
	</svelte:fragment>
	<HeaderUtilities>
		<HeaderActionLink
			icon={LogoGithub}
			href="https://github.com/dpkg-i-foo-deb/libre-asi"
			target="_blank"
		/>
		{#if $setup}
			{#if $session.active}
				<HeaderGlobalAction
					aria-label="Settings"
					icon={SettingsAdjust}
					on:click={function () {
						goto('/settings');
					}}
				/>
			{/if}

			<HeaderAction
				bind:isOpen={isUserMenuOpen}
				icon={UserAvatarFilledAlt}
				closeIcon={UserAvatarFilledAlt}
			>
				<HeaderPanelLinks>
					{#if !$session.active}
						<HeaderPanelDivider>{$LL.navMenu.NOT_LOGGED_IN()}</HeaderPanelDivider>
						<HeaderPanelLink
							href="/login"
							on:click={() => {
								isUserMenuOpen = false;
							}}>{$LL.navMenu.LOGIN()}</HeaderPanelLink
						>
					{:else if $session.active}
						<HeaderPanelDivider>{$LL.navMenu.SESSION_ACTIVE()}</HeaderPanelDivider>
						<HeaderPanelLink on:click={handleSignOut}>{$LL.navMenu.SIGN_OUT()}</HeaderPanelLink>
					{/if}
				</HeaderPanelLinks>
			</HeaderAction>
		{/if}
	</HeaderUtilities>
</Header>

{#if $setup && $page.url.pathname != '/setup' && $page.url.pathname != '/login'}
	<SideNav bind:isOpen={isSideNavOpen}>
		<SideNavItems>
			<SideNavLink
				text={$LL.navMenu.WELCOME_PAGE()}
				href="/"
				isSelected={$page.url.pathname == '/'}
			/>

			{#if $session.active}
				<SideNavLink
					text={$LL.navMenu.HOME()}
					href="/home"
					isSelected={$page.url.pathname == '/home'}
				/>
				<SideNavMenu
					text={$LL.navMenu.MANAGEMENT()}
					expanded={$page.url.pathname.toString().includes('/management/')}
				>
					{#if $session.role == SessionRole.Admin}
						<SideNavMenuItem
							href="/management/administrators"
							text="Administradores"
							isSelected={$page.url.pathname == '/management/administrators'}
						/>

						<SideNavMenuItem
							text="Entrevistadores"
							href="/management/interviewers"
							isSelected={$page.url.pathname == '/management/interviewers'}
						/>
					{/if}

					{#if $session.role == SessionRole.Admin || $session.role == SessionRole.Interviewer}
						<SideNavMenuItem
							href="/management/patients"
							text="Pacientes"
							isSelected={$page.url.pathname == '/management/patients'}
						/>
					{/if}

					{#if $session.role == SessionRole.Admin || $session.role == SessionRole.Interviewer}
						<SideNavMenuItem
							href="/management/interviews"
							text="Entrevistas"
							isSelected={$page.url.pathname == '/management/interviews'}
						/>
					{/if}
				</SideNavMenu>
			{/if}
		</SideNavItems>
	</SideNav>
{/if}

<Content>
	<Grid padding={false} style="padding-left:0;">
		<Row>
			<Column style="padding-left:0;">
				{#if !canRender}
					{#await checkSetup()}
						<ProgressBar helperText={$LL.navMenu.LOADING()} />
					{:then}
						{#if canRender}
							<slot />
						{/if}
					{/await}
				{:else}
					<slot />
				{/if}

				<div class="notification">
					<Notification />
				</div>
			</Column>
		</Row>
	</Grid>
</Content>

<style>
	.notification {
		position: fixed;
		bottom: 0;
		right: 0;
		margin: 10px;
	}
</style>
