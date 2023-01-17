<script lang="ts">
	import 'carbon-components-svelte/css/all.css';
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
		SideNavLink,
		SkipToContent,
		Content,
		Grid,
		Row,
		Column,
		HeaderActionLink,
		ToastNotification
	} from 'carbon-components-svelte';
	import SettingsAdjust from 'carbon-icons-svelte/lib/SettingsAdjust.svelte';
	import UserAvatarFilledAlt from 'carbon-icons-svelte/lib/UserAvatarFilledAlt.svelte';
	import LogoGithub from 'carbon-icons-svelte/lib/LogoGithub.svelte';
	import { session } from '$lib/stores/userStore';
	import type { Unsubscriber } from 'svelte/store';
	import { onDestroy, onMount } from 'svelte';
	import { apiUrl, signOut } from '$lib/api/constants';
	import { goto } from '$app/navigation';
	import { loggedInCorrectly } from '$lib/stores/loginStore';
	import { page } from '$app/stores';

	let isSideNavOpen = false;
	let isUserMenuOpen = false;
	let subscription: Unsubscriber;
	let activeSession = false;
	let logOutError = false;
	let loggedOutCorrectly = false;

	onMount(async () => {
		subscription = session.subscribe((value) => {
			if (value == 'true') {
				activeSession = true;
			} else {
				activeSession = false;
			}
		});
	});

	onDestroy(() => {
		subscription;
	});

	async function logOut() {
		logOutError = false;
		loggedOutCorrectly = false;
		let result: Response;
		try {
			result = await fetch(apiUrl + signOut, {
				method: 'POST',
				credentials: 'include',
				mode: 'cors'
			});
		} catch (e) {
			logOutError = true;
			console.log(e);
			return;
		} finally {
			isUserMenuOpen = false;
		}

		if (result.ok || result.status == 401) {
			session.set('false');
			goto('/');
			loggedOutCorrectly = true;
		} else {
			logOutError = true;
		}
	}
</script>

<Header persistentHamburgerMenu={true} bind:isSideNavOpen company="Libre-ASI">
	<svelte:fragment slot="skip-to-content">
		<SkipToContent />
	</svelte:fragment>
	<HeaderUtilities>
		<HeaderActionLink
			icon={LogoGithub}
			href="https://github.com/dpkg-i-foo-deb/libre-asi"
			target="_blank"
		/>
		{#if activeSession}
			<HeaderGlobalAction aria-label="Settings" icon={SettingsAdjust} />
		{/if}

		<HeaderAction
			bind:isOpen={isUserMenuOpen}
			icon={UserAvatarFilledAlt}
			closeIcon={UserAvatarFilledAlt}
		>
			<HeaderPanelLinks>
				{#if !activeSession}
					<HeaderPanelDivider>Not Logged in</HeaderPanelDivider>
					<HeaderPanelLink
						href="/login"
						on:click={() => {
							isUserMenuOpen = false;
						}}>Log In</HeaderPanelLink
					>
				{:else if activeSession}
					<HeaderPanelDivider>Session Active</HeaderPanelDivider>
					<HeaderPanelLink on:click={logOut}>Sign Out</HeaderPanelLink>
				{/if}
			</HeaderPanelLinks>
		</HeaderAction>
	</HeaderUtilities>
</Header>

<SideNav bind:isOpen={isSideNavOpen}>
	<SideNavItems>
		<SideNavLink text="Welcome Page" href="/" isSelected={$page.url.pathname == '/'} />
		<SideNavLink text="Home" href="/home" isSelected={$page.url.pathname == '/home'} />
		<SideNavLink text="Dashboard" href="/" isSelected={$page.url.pathname == '/dashboard'} />
	</SideNavItems>
</SideNav>

<Content>
	<Grid>
		<Row>
			<Column>
				<slot />
				{#if logOutError}
					<div class="notification">
						<ToastNotification
							title="Sign Out Error"
							subtitle="Try again later"
							caption="If the error persists, contact your administrator"
						/>
					</div>
				{/if}

				{#if loggedOutCorrectly}
					<div class="notification">
						<ToastNotification
							kind="info"
							title="Signed Out Correctly"
							subtitle="Thanks for trying out Libre-ASI!"
							caption={new Date().toLocaleString()}
						/>
					</div>
				{/if}

				{#if $loggedInCorrectly}
					<div class="notification">
						<ToastNotification
							kind="success"
							title="Logged In"
							subtitle="Welcome to Libre-ASI"
							caption={new Date().toLocaleString()}
						/>
					</div>
				{/if}
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
