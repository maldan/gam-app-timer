<template>
  <div :class="$style.grid">
    <div :class="$style.block">
      <Button text="5" />
    </div>

    <div :class="$style.block">
      <Button text="Hi" />
      <Button text="Hi" color="gray" />

      <Checkbox v-model="sas" />
      <Checkbox v-model="sas" />
      <Checkbox v-model="sas" />

      <Toggle :list="['1', '2']" v-model="sasx" />

      <Input />
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Button,
  Checkbox,
  Toggle,
  Input,
} from "../gam-lib-ui/vue/component/ui";
import { ref } from "vue";

const sas = ref(false);
const sasx = ref("1");
</script>

<style module lang="scss">
@import "../gam-lib-ui/vue/vars";

.grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.block {
  padding: 20px;
  background-color: #191a1d;
  border-radius: $radius-3;
  display: flex;
  flex-direction: column;

  /*button {
    background-color: transparent;
    color: $color-white-060;
    outline: none;
    padding: calc($size-1 - 2px) calc($size-2 - 2px);
    margin: 2px;
    border: 2px solid transparent;
    border-radius: $radius-1;
    transition: border-color 0.2s ease;
    cursor: pointer;

    &:hover {
      background-color: #212225;
    }
  }*/

  .button_2 {
    background: $color-main;
    color: #fff;
    font-style: normal;
    font-weight: 400;
    padding-left: $space-3;
    padding-right: $space-3;
    box-shadow: inset 0 0 0 1px $color-white-010;

    &:hover {
      background-color: darken($color-main, 10%);
    }
  }
}
</style>
