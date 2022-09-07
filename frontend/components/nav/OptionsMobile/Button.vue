<template>
    <div id="container-button" class="flex-vertical flex-center-align flex-space-between clickable">
        <div class="s-part"></div>
        <div class="s-part"></div>
        <div class="s-part"></div>
    </div>
</template>

<style scoped>

    #container-button
    {
        height: calc(var(--nav-height) * 0.6);
        min-height: calc(var(--nav-min-height) * 0.6);
        width: calc(var(--nav-height) * 0.6 + 3vw);
        min-width: calc(var(--nav-min-height) * 0.6 + 3vw);
        
    }

    #container-button:hover>*
    {
        background: var(--primary-color);
    }

    .s-part
    {
        width: 100%;
        height: 22%;
        background: var(--text-color);
        border-radius: 5px;
        transition: 150ms ease-in-out;
    }

</style>