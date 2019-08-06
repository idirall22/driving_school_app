<template>
  <div class="">
    <b-form-group id="searchStudents"
      label="Cherché des étudiants:" label-for="searchStudents">
      <b-form-input
        id="searchStudents"
        v-model="studentLastName"
        type="text"
        placeholder="Nom de l'étudiant">
      </b-form-input>
    </b-form-group>

    <b-button
      @click="searchStudent()"
      variant="primary">
      Cherché
    </b-button>

      <div v-if="found">
        <div class="col mt-3 p-0">
          <ul class="list-group p-0 m-0">
            <li v-for="student in studentsFound" :key="student.id"
            class="list-group-item d-flex justify-content-between align-items-center">
            <p class="p-0 m-0">
              {{student.last_name_fr | capitalize}}
              {{student.first_name_fr | capitalize}}
            </p>
            <button
              class="btn btn-primary m-0"
              @click="addStudent(student)"
              type="button"
              name="button"
              >Ajouté
            </button>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import Student from '../service/student.js'
export default {
  name: "searchStudent",
  data: () => ({
    studentsFound: [],
    studentLastName: "mak",
    found: false,
  }),
  methods: {
    addStudent:function(student){
      this.studentsFound = [];
      this.found = false;
      this.studentLastName="";
      this.$emit("studentAdded", student);
    },
    searchStudent:function(){
      //Check if student name length is greather then 2
      if(this.studentLastName.length > 2){
        window.backend.Service.GetStudents(
          this.studentLastName, "", 10, 0)
          .then(
            data=>{
              // Convert data to student object
              for (var i = 0; i < data["students"].length; i++) {
                let student = new Student(data["students"][i], null);
                this.studentsFound.push(student)
              }
              this.found = true;
        })
      }else{
        this.studentsFound = [];
        this.found = false;
      }
    },
  }
}
</script>
<style scoped>
</style>