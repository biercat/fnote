<template>
  <div
    class="footer_shadow flex flex-col justify-evenly items-center p-10 dark:text-dtc dark_bg_gray shadow-xl"
  >
    <div>
      Copyright © {{ year }} - Designed by
      <a
        class="decoration-none text-#000 font-bold hover:underline dark_text_white"
        href="https://github.com/chenmingyong0423/fnote"
        target="_blank"
        >Fnote</a
      >
    </div>
    <div class="flex items-center lt-lg:flex-col">
      <span v-for="(item, index) in records" :key="index" v-html="item"> </span>
    </div>
    <div></div>
  </div>
</template>

<script lang="ts" setup>
import { useConfigStore } from "~/store/config";

const configStore = useConfigStore();

const date = new Date(configStore.website_info.website_runtime * 1000);
const year = date.getFullYear();

const records = configStore.website_info.website_records;
</script>

<style scoped></style>
