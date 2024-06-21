<template>
  <div
    style="
      margin: 100px;
      font-size: 30px;
      font-family: '楷体';
      font-weight: bold;
    "
  >
    son1
    {{ testObj }}
  </div>
</template>
<script setup>
import { getMyTest } from "@/api/mytest.js";
import { onMounted, ref } from "vue";
const testObj = ref("");
const getMyTestData = async () => {
  const res = await getMyTest();
  testObj.value = res;
  console.log(res);
};
onMounted(() => {
  getMyTestData();
});
</script>
