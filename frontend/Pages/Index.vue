<template>
<div class="my-5 text-center">
    <div class="text-center">
        <OfcLogo />
    </div>

    <h1 class="mt-3">Starter kit success!</h1>
    <p>If you are looking at this page in the browser, then you're probably good to go! Start building your app!</p>

    <p>
        <a href="https://www.openfunctioncomputers.com/" target="_blank">Open Function Computers, LLC</a> |
        <a href="https://github.com/open-function-computers-llc/go-inertia-starter" target="_blank">
            <GithubLogo />
            Fork this on Github
        </a>
    </p>
</div>
</template>

<script setup>
import GithubLogo from "@/Components/GithubLogo.vue";
import OfcLogo from "@/Components/OfcLogo.vue";
import Layout from "@/Layouts/Standard.vue";

defineOptions({ layout: Layout });
</script>

<style lang="scss" scoped>
@use "@/scss/variables" as *;

a {
    text-decoration: none;
    color: $c-gray;
}
</style>
