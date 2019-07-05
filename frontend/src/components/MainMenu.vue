<template>
  <div id="mainMenu">
    <nav class="d-none d-md-block border-right">
      <div class="sidebar-sticky">
        <ul class="nav flex-column">
          <div class="drivingSchoolName">
            <h5>Driving School</h5>
          </div>
          <!-- TODO make better solution -->
          <!-- Students list -->
          <li class="nav-item" v-on:click="setCurrentPage(0)"
            v-bind:class="['',(currentPage === 0)?'active':'']">
            <router-link to="/" class="nav-link">
              Students
            </router-link>
          </li>

            <!-- Add new student -->
          <li class="nav-item" v-on:click="setCurrentPage(1)"
            v-bind:class="['',(currentPage === 1)?'active':'']">
            <router-link to="/add-student" class="nav-link">
              Add Student
            </router-link>
          </li>

            <!-- Add exam list -->
          <li class="nav-item" v-on:click="setCurrentPage(2)"
            v-bind:class="['',(currentPage === 2)?'active':'']">
            <router-link to="/add-exam-list" class="nav-link">
              Add Exam List
            </router-link>
          </li>

            <!-- exam lists -->
          <li class="nav-item" v-on:click="setCurrentPage(3)"
            v-bind:class="['',(currentPage === 3)?'active':'']">
            <router-link to="/exam-lists" class="nav-link">
              Exam Lists
            </router-link>
          </li>
        </ul>
      </div>
    </nav>
  </div>
</template>

<script>

export default {
  name: "mainMenu",
  data: () => ({
    currentPage: 0,
    items:[],
    pages: [
      {"title": "Students","route": ""},
      {"title": "Add Student", "route": "add-student"},
      {"title": "Add Exam List", "route": "add-exam-list"},
      {"title": "Exam Lists", "route": "exam-lists"}
    ],
  }),
  methods:{
    setCurrentPage: function(index){
      this.currentPage = index;
    }
  }
}
</script>

<style scoped>
nav{
  height: 100%;
}
.drivingSchoolName{
  margin: 0px;
  padding: 25px;
  color: black;
}
.nav-link{
  padding-top: 20px;
  padding-bottom: 20px;
}

a{
  margin: 0px;
  color: black;
}
a:hover{
  color: #0486ff;
  background-color: #e3e3e3;
}
.active{
  border-right: 4px solid #0486ff;
  background-color: #e3e3e3;
  color: #0486ff;
}

</style>