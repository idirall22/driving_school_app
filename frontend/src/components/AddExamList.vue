<template>
  <div id="addExamList">

    <!-- title -->
    <div class="row">
      <Header initTitle="Add Exam List"></Header>

      <!-- Message  -->
       <Message
        v-on:close-alert="closeAlert"
        :create.sync="examListCreated"
        :errorCreate.sync="errorCreateExamList"
        subject="Exam List">
      </Message>
    </div>

    <!-- exam list form -->
    <ExamListForm :students="students" ref="examListForm"></ExamListForm>

    <!-- array students -->
    <div class="table-responsive mt-5">
      <table class="table table-striped table-sm">
        <thead>
          <tr>
            <th>N°</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Next Exam</th>
            <th>Details</th>
            <th></th>
          </tr>

        </thead>
        <tbody>
          <tr v-for="(student, index) in students" :key="student.key">
            <td>{{student.id}}</td>
            <td>{{student.first_name}}</td>
            <td>{{student.last_name}}</td>
            <td>{{student.next_exam}}</td>
            <td><a class="btn btn-warning" href="#">Infos</a></td>
            <td><a class="btn btn-danger" @click="deleteStudent(index)">X</a></td>
          </tr>
        </tbody>
      </table>
      <button @click.prevent="addExamList()"
              class="btn btn-primary btn-lg btn-block"
              type="submit">Add Exam List</button>
    </div>
  </div>
</template>
<script>

import Header from './parts/Header';
import Message from './parts/Message';
import ExamListForm from './parts/ExamListForm';
import moment from 'moment'

export default {
  name: "addExamList",
  components: {
    ExamListForm,
    Message,
    Header
  },
  data: () => ({
    errorCreateExamList: null,
    examListCreated: false,
    studentLastName: "",
    examinerName: "",
    examDate: null,
    found: false,
    students: []
  }),
  methods:{

    // Add exam list
    addExamList: function(){
      window.backend.Service.CreateExamList(
        moment(this.$refs.examListForm.examDate).format(),
        this.$refs.examListForm.examinerName,
        this.$refs.examListForm .students)

        .then(
          data=>{this.data = data},
          err=>{this.errorCreateExamList= err}
        )

      if(this.errorCreateExamList != null){
        this.examListCreated = false;
      }else{
        this.examListCreated = true;
        // this.$refs.examListForm.resetForm();
      }
    },

    // Delete student from exam list
    deleteStudent: function(studentID){
      this.students.splice(studentID, 1)
    },
    closeAlert: function(){
      this.errorCreateExamList = null;
      this.examListCreated =  false;
    }
  },
}
</script>
<style scoped>
</style>
