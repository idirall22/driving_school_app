<template>
  <div id="examListDetails">
    <div class="mb-5 mt-3">
      <h1>Exam List</h1>
    </div>

    <div class="table-responsive">
      <table class="table table-striped table-sm">
        <thead>
          <tr>
            <th>Order</th>
            <th>fileNumber</th>
            <th>Full Name</th>
            <th>Birthday</th>
            <th>Cat</th>
            <th>Exam Name</th>
            <th>Results</th>
          </tr>

        </thead>
        <tbody>
          <!-- <tr v-for="(exam, index) in examLists" :key="exam.id">
            <td>{{index}}</td>
          </tr> -->
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "examListDetails",
  data: () => ({

  }),
  methods:{

  }
}
</script>
<style scoped>
</style>