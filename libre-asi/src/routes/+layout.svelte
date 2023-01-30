<script lang="ts">
	import 'carbon-components-svelte/css/all.css';
	import Notification from '../components/Notification.svelte';
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
		SideNavMenu
	} from 'carbon-components-svelte';
	import SettingsAdjust from 'carbon-icons-svelte/lib/SettingsAdjust.svelte';
	import UserAvatarFilledAlt from 'carbon-icons-svelte/lib/UserAvatarFilledAlt.svelte';
	import LogoGithub from 'carbon-icons-svelte/lib/LogoGithub.svelte';
	import session from '$lib/stores/userStore';
	import { page } from '$app/stores';
	import { SessionRole } from '$lib/models/Session';
	import { notifications } from '$lib/stores/notificationStore';
	import { goto } from '$app/navigation';
	import { sendSuccess } from '$lib/util/notifications';

	let isSideNavOpen = false;
	let isUserMenuOpen = false;

	async function handleSignOut() {
		try {
			await fetch('api/sign-out', { method: 'POST', credentials: 'include' });

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
		{#if $session.active}
			<HeaderGlobalAction aria-label="Settings" icon={SettingsAdjust} />
		{/if}

		<HeaderAction
			bind:isOpen={isUserMenuOpen}
			icon={UserAvatarFilledAlt}
			closeIcon={UserAvatarFilledAlt}
		>
			<HeaderPanelLinks>
				{#if !$session.active}
					<HeaderPanelDivider>Not Logged in</HeaderPanelDivider>
					<HeaderPanelLink
						href="/login"
						on:click={() => {
							isUserMenuOpen = false;
						}}>Log In</HeaderPanelLink
					>
				{:else if $session.active}
					<HeaderPanelDivider>Session Active</HeaderPanelDivider>
					<HeaderPanelLink on:click={handleSignOut}>Sign Out</HeaderPanelLink>
				{/if}
			</HeaderPanelLinks>
		</HeaderAction>
	</HeaderUtilities>
</Header>

<SideNav bind:isOpen={isSideNavOpen}>
	<SideNavItems>
		<SideNavLink text="Welcome Page" href="/" isSelected={$page.url.pathname == '/'} />
		<SideNavLink text="Home" href="/home" isSelected={$page.url.pathname == '/home'} />
		<SideNavMenu
			text="Management"
			expanded={$page.url.pathname.toString().includes('/management/')}
		>
			{#if $session.role == SessionRole.Admin}
				<SideNavMenuItem
					href="/management/admin"
					text="Administrators"
					isSelected={$page.url.pathname == '/management/admin'}
				/>

				<SideNavMenuItem text="Interviewers" />
			{/if}
		</SideNavMenu>
	</SideNavItems>
</SideNav>

<Content>
	<Grid>
		<Row>
			<Column>
				<slot />

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
