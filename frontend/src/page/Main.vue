<template>
  <div :class="$style.main">
    <Block>
      <div v-if="!isChange" :class="$style.time">12:30:43</div>
      <Input v-if="isChange" v-model="time" />

      <Row size="2" style="margin-top: 10px">
        <Button v-if="!isChange" @click="isChange = !isChange" text="Change" />
        <Button v-if="isChange" @click="sas" text="Ok" />
        <Button text="Stop" />
      </Row>
    </Block>
  </div>
</template>

<script setup lang="ts">
import {
  Button,
  Checkbox,
  Toggle,
  Input,
  Block,
  Row,
} from "../gam-lib-ui/vue/component/ui";
import { ref } from "vue";
import dayjs from "dayjs";
import { TimeHelper } from "@/gam-lib-ui/util/TimeHelper";

const time = ref("00:00:00");
const isChange = ref(false);

function sas() {
  isChange.value = false;
  console.log(TimeHelper.HMStoSeconds(time.value));
}
</script>

<style module lang="scss">
@import "../gam-lib-ui/vue/vars";

.main {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;

  .time {
    font-size: 24px;
    text-align: center;
  }
}
</style>
