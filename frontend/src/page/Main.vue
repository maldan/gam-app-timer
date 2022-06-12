<template>
  <div :class="$style.main">
    <Block> Hello </Block>
  </div>
</template>

<script setup lang="ts">
import {
  Button,
  Checkbox,
  Toggle,
  Input,
  Block,
} from "../gam-lib-ui/vue/component/ui";
import { ref } from "vue";

const sas = ref(false);
const sasx = ref("1");
</script>

<style module lang="scss">
@import "../gam-lib-ui/vue/vars";

.main {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}
</style>
