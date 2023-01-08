<script lang="ts">
    import {t,locale,locales} from "../util/i18n"
    let color = '#1e90ff';
</script>

<main>
    <div class="footer-background" style="--color:{color}">
        <div class="text-container">{$t("footer.connected")}</div>
    </div>
</main>

<style>

    .footer-background{
        background-color:var(--color);
        width: 100%;
        height: 22px;
        padding-bottom: 0%;
    }

    .text-container{
        display: inline-block;
        margin: 0 auto;
        padding-left: 12px;
        color: white;
    }

</style>