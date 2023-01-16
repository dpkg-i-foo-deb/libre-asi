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

	let isSideNavOpen = false;
	let isOpen1 = false;
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
		<HeaderGlobalAction aria-label="Settings" icon={SettingsAdjust} />

		<HeaderAction bind:isOpen={isOpen1} icon={UserAvatarFilledAlt} closeIcon={UserAvatarFilledAlt}>
			<HeaderPanelLinks>
				<HeaderPanelDivider>Switcher subject 1</HeaderPanelDivider>
				<HeaderPanelLink>Switcher item 1</HeaderPanelLink>
				<HeaderPanelLink>Switcher item 2</HeaderPanelLink>
				<HeaderPanelLink>Switcher item 3</HeaderPanelLink>
				<HeaderPanelLink>Switcher item 4</HeaderPanelLink>
				<HeaderPanelDivider>Switcher subject 2</HeaderPanelDivider>
				<HeaderPanelLink>Switcher item 1</HeaderPanelLink>
				<HeaderPanelLink>Switcher item 2</HeaderPanelLink>
				<HeaderPanelDivider>Switcher subject 3</HeaderPanelDivider>
				<HeaderPanelLink>Switcher item 1</HeaderPanelLink>
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
