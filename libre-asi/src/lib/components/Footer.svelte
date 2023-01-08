<script lang="ts">
  import { onMount } from "svelte";
  import { ping } from "../api/ping";
  import { t } from "../util/i18n";

  let color = "#1e90ff";
  let poller;
  let connected: Boolean;

  const setUpPing = () => {
    pingApi();
    poller = setInterval(pingApi, 15000);
  };

  const pingApi = async () => {
    connected = await ping();

    if (!connected) {
      color = "#e32636";
    } else {
      color = "#1e90ff";
    }
  };

  onMount(setUpPing);
</script>

<main>
  <div class="footer-background" style="--color:{color}">
    {#if connected}
      <div class="text-container">{$t("footer.connected")}</div>
    {/if}

    {#if !connected}
      <div class="text-container">{$t("footer.disconnected")}</div>
    {/if}
  </div>
</main>

<style>
  .footer-background {
    background-color: var(--color);
    width: 100%;
    height: 22px;
    padding-bottom: 0%;
  }

  .text-container {
    display: inline-block;
    margin: 0 auto;
    padding-left: 12px;
    color: white;
  }
</style>
