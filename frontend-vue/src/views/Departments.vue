<template>
  <div class="departments">
    <h1>Departments</h1>
    <DepartmentCard
      v-for="department in departments"
      :key="department.id"
      :department="department"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useQuery, useResult } from "@vue/apollo-composable";
import { gql } from "@apollo/client";

import { Department } from "@/types";

import DepartmentCard from "@/components/DepartmentCard.vue";

const allDepartmentsQuery = gql`
  query allDepartments {
    departments {
      id
      name
    }
  }
`;

export default defineComponent({
  name: "Departments",
  components: { DepartmentCard },
  setup() {
    const { result: departmentsResult } = useQuery(allDepartmentsQuery);
    const departments = useResult(
      departmentsResult,
      [] as Department[],
      (data) => data.departments as Department[]
    );

    return {
      departments,
    };
  },
});
</script>
