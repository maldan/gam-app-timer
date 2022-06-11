<template>
  <div :class="$style.grid">
    <div :class="$style.block">
      Nado salo est
      <button>s</button>
      <button>s</button>
      <button>s</button>

      <Button />
    </div>

    <div :class="$style.block">
      Nado salo est
      <button>s</button>
      <button>s</button>
      <button>s</button>

      <Button />

      <Checkbox v-model="sas" />
      <Checkbox v-model="sas" />
      <Checkbox v-model="sas" />

      <Toggle :list="['1', '2']" v-model="sasx" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { Button, Checkbox, Toggle } from "../gam-lib-ui/vue/component/ui";
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

  button {
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
  }

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
