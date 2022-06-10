<template>
  <div :class="$style.block">
    Nado salo est
    <button>s</button>
    <button>s</button>
    <button>s</button>
    <button :class="$style.button_2">s</button>
  </div>
</template>

<script setup lang="ts"></script>

<style module lang="scss">
@import "../vars.scss";

.block {
  padding: 20px;
  background-color: #191a1d;
  border-radius: var(--radius-3);
  display: flex;
  flex-direction: column;

  button {
    background-color: transparent;
    color: var(--color-white-060);
    outline: none;
    padding: calc(var(--size-1) - 2px) calc(var(--size-2) - 2px);
    margin: 2px;
    border: 2px solid transparent;
    border-radius: var(--radius-1);
    transition: border-color 0.2s ease;
    cursor: pointer;

    &:hover {
      background-color: #212225;
    }
  }

  .button_2 {
    background: var(--color-main);
    color: #fff;
    font-style: normal;
    font-weight: 400;
    padding-left: var(--space-3);
    padding-right: var(--space-3);
    box-shadow: inset 0 0 0 1px var(--color-white-010);

    &:hover {
      background-color: darken($color-main, 10%);
    }
  }
}
</style>
