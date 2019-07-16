<template>
  <div v-if="loaded" id="examListDetails">
    <p>students already in exam: {{examList.students_exams}}</p>
    <p>new students: {{changeStudentsExams}}</p>
    <div class="row">
      <Header initTitle="Exam List"></Header>
      <!-- Message  -->
       <Message
        v-on:close-alert="closeAlert"
        :create.sync="examListUpdated"
        :errorCreate.sync="errorUpdateExamList"
        subject="subjectMessage">
      </Message>
    </div>
    <div v-if="loaded" class="">
      <!-- exam list form -->
      <ExamListForm
      :examinerName="examList.examiner"
      :examDate="examDate"
      :students="changeStudentsExams"
      ref="examListForm">
    </ExamListForm>
    </div>


    <hr>
    <div class="table-responsive">
      <table class="table table-striped table-sm">
        <thead>
          <tr>
            <th class="align-middle text-center">Order</th>
            <th class="align-middle text-center">fileNumber</th>
            <th class="align-middle text-center">Full Name</th>
            <th class="align-middle text-center">Birthday</th>
            <th class="align-middle text-center">Cat</th>
            <th class="align-middle text-center">Exam Name</th>
            <th class="align-middle text-center">Results</th>
            <th class="align-middle text-center">Remove</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(exam, index) in examList.students_exams" :key="exam.id">
            <td class="align-middle text-center">{{index+1}}</td>
            <td class="align-middle text-center">{{exam.student.file_number}}</td>
            <td class="align-middle text-center">{{exam.student.first_name | capitalize}} {{exam.student.last_name | capitalize}}</td>
            <td class="align-middle text-center">{{exam.student.birthday | moment2}}</td>
            <td class="align-middle text-center">B</td>
            <td class="align-middle text-center">{{exam.exam}}</td>
            <td class="align-middle text-center">
              <input v-model="exam.status" type="checkbox"> </td>
            <td class="align-middle text-center">
              <button @click.prevent="removeStudent(exam)"
                type="button" class="form-control close"
                data-dismiss="alert" aria-label="Close">
               <span aria-hidden="true">&times;</span>
             </button>
           </td>
          </tr>

          <tr v-for="(student, index) in studentsAdded" :key="student.id">
            <td class="align-middle text-center">{{index+1}}</td>
            <td class="align-middle text-center">{{student.file_number}}</td>
            <td class="align-middle text-center">{{student.first_name | capitalize}} {{student.last_name | capitalize}}</td>
            <td class="align-middle text-center">{{student.birthday | moment2}}</td>
            <td class="align-middle text-center">B</td>
            <td class="align-middle text-center">{{student.next_exam}}</td>
            <td class="align-middle text-center">
              <input type="checkbox"> </td>
            <td class="align-middle text-center">
              <button @click.prevent="removeStudentAdded(student.id)"
                type="button" class="form-control close"
                data-dismiss="alert" aria-label="Close">
               <span aria-hidden="true">&times;</span>
             </button>
           </td>
          </tr>

        </tbody>
      </table>
    </div>
    <!-- button add student -->
    <button @click.prevent="editexamList()"
      class="btn btn-primary btn-lg btn-block"
      type="submit">Edit Exam List
    </button>
  </div>
</template>

<script>
import Header from './parts/Header';
import Message from './parts/Message';
import ExamListForm from './parts/ExamListForm';
import moment from 'moment'

export default {
  components: {
    Header,
    Message,
    ExamListForm
  },

  mounted() {
    this.getExamList();
  },
  name: "ExamListDetails",
  data: () => ({
    subjectMessage: "",
    examListUpdated: false,
    errorUpdateExamList: null,
    errorGetExamList: null,
    id: 0,
    examList: {},
    loaded: false,
    examDate: null,
    studentsAdded: [],
  }),
  watch:{
    studentsAdded: function(){
      // Always check if the last student added is already
      // in the list if yes delete it
      if(this.studentsAdded.length > 0){
        if(this.examList.students_exams){
          for (var i = 0; i < this.examList.students_exams.length; i++) {
            if(this.examList.students_exams[i].student.id ==
              this.studentsAdded[this.studentsAdded.length -1 ].id){
              this.removeStudentAdded(this.studentsAdded[this.studentsAdded.length -1 ].id)
              return
            }
          }
        }
      }
    }
  },
  computed:{
    changeStudentsExams:{
      get: function(){return this.studentsAdded;},
      set: function(value){this.studentsAdded = value;}
    }
  },
  methods:{
    getExamList: function(){
      // Get id from url
      this.id = this.$route.params.id;
      if(this.id != 0){
        // Make request to backend
        window.backend.Service.GetExamList(this.id)
        .then(
          data=>{
            this.examList = data;
            this.examDate = moment(this.examList.date_exam).format();
            this.loaded = true;
          },
          err=>{this.errorGetExamList = err;}
        );
      }
    },
    removeStudentAdded: function(studentID){
      for (var i = 0; i < this.studentsAdded.length; i++) {
        if(this.studentsAdded[i].id == studentID){
          this.studentsAdded.splice(i, 1)
        }
      }
    },
    removeStudent: function(exam){
      for (var i = 0; i < this.examList.students_exams.length; i++) {
        if(this.examList.students_exams[i].id == exam.id){
          this.examList.students_exams.splice(i, 1);
        }
      }
    },

    editexamList: function(){
      // Update Exam list
      window.backend.Service.UpdateExamList(
        this.examList.id,
        moment(this.$refs.examListForm.changeExamDate).format(),
        this.$refs.examListForm.changeExaminerName,
        [],
        this.studentsAdded,
      ).then(
        data=>{this.data= data},
        err=>{this.errorUpdateExamList = err},
      );

      if(this.errorUpdateExamList != null){
        this.examListUpdated = false;
      }else{
        this.examListUpdated = true;
        this.errorUpdateExamList = null;
      }

    },
    closeAlert:function(){
      this.examListUpdated = false;
      this.errorUpdateExamList = null;
    }
  }
}
</script>
<style scoped>
</style>