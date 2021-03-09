<template>
  <div class="departments">
    <img alt="Vue logo" src="../assets/logo.png" />
    <h1>Departments</h1>
    <ul>
      <li v-for="department in departments" :key="department.id">
        <DepartmentCard :department="department" />
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

import { Department } from "@/types";

import DepartmentCard from "@/components/DepartmentCard.vue";

const _departments: Department[] = [
  { id: 1, name: "Dep 1" },
  { id: 2, name: "Dep 2" },
  { id: 3, name: "Dep 3" },
];

export default defineComponent({
  name: "Departments",
  components: { DepartmentCard },
  setup() {
    const departments = _departments;

    return {
      departments,
    };
  },
});
</script>
