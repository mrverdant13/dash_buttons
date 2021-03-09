<template>
  <p>
    {{ department.name }}
  </p>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";

import { Department } from "@/types";

export default defineComponent({
  name: "DepartmentCard",
  props: {
    department: {
      type: Object as PropType<Department>,
      required: true,
    },
  },
});
</script>


<style>
</style>