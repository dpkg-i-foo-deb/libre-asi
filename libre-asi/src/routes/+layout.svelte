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
		HeaderActionLink
	} from 'carbon-components-svelte';
	import SettingsAdjust from 'carbon-icons-svelte/lib/SettingsAdjust.svelte';
	import UserAvatarFilledAlt from 'carbon-icons-svelte/lib/UserAvatarFilledAlt.svelte';
	import LogoGithub from 'carbon-icons-svelte/lib/LogoGithub.svelte';
	import { session } from '$lib/stores/userStore';
	import type { Unsubscriber } from 'svelte/store';
	import { onDestroy, onMount } from 'svelte';

	let isSideNavOpen = false;
	let isUserMenuOpen = false;
	let subscription: Unsubscriber;
	let activeSession = false;

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

	function logOut() {
		console.log('uwu');
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
					<HeaderPanelLink>Log In</HeaderPanelLink>
				{:else if activeSession}
					<HeaderPanelDivider>Session Active</HeaderPanelDivider>
					<HeaderPanelLink on:click={logOut}>Sign Out</HeaderPanelLink>
				{:else}{/if}
			</HeaderPanelLinks>
		</HeaderAction>
	</HeaderUtilities>
</Header>

<SideNav bind:isOpen={isSideNavOpen}>
	<SideNavItems>
		<SideNavLink text="Home" href="/home" />
		<SideNavLink text="Dashboard" href="/" />
	</SideNavItems>
</SideNav>

<Content>
	<Grid>
		<Row>
			<Column>
				<slot />
			</Column>
		</Row>
	</Grid>
</Content>
