<template>
  <div id="examListDetails">
    <div class="row">
      <div class="col mb-5 mt-3">
        <h1>Exam List</h1>
        <button @click="switchEditMode" class="btn btn-warning">Edit mode</button>
      </div>
      <div v-if="examListUpdated"
        class="col mb-5 mt-3 mr-3 alert alert-success" role="alert">
        <button @click.prevent="closeAlert()" type="button" class="close"
            data-dismiss="alert" aria-label="Close">
           <span aria-hidden="true">&times;</span>
         </button>
        <h4 class="alert-heading">Exam List Updated!</h4>
        <hr>
        <p>The exam list was updated Successfuly</p>
      </div>
      <div v-if="!examListUpdated && errorUpdateExamList != null" class="col mb-5 mt-3 mr-3 alert alert-warning" role="alert">
        <button @click.prevent="closeAlert()" type="button" class="close"
            data-dismiss="alert" aria-label="Close">
           <span aria-hidden="true">&times;</span>
         </button>
        <h4 class="alert-heading">Could not update the exam list!</h4>
        <hr>
        <p>The exam list was not updated</p>
      </div>
    </div>
    <!-- block one -->
    <div class="row">
      <!-- examiner name -->
      <div class="col-6">
        <label for="examinerName">Examiner</label>
        <input v-model="examList.examiner" type="text" class="form-control" id="examinerName"
        placeholder="Search">
      </div>
      <!-- date picker -->
      <div class="col-6">
        <label for="datepicker">Exam date</label>
        <vue-bootstrap-datetimepicker v-model="examDate"
        id="datepicker" :config="options"></vue-bootstrap-datetimepicker>
      </div>
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
              <input v-model="exam.status" type="checkbox" :disabled="!editMode?true: false"> </td>
            <td class="align-middle text-center">
              <button @click.prevent="removeStudent(exam)"
                type="button" class="form-control close"
                data-dismiss="alert" aria-label="Close"
                :disabled="!editMode?true: false">
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
      type="submit" :disabled="!editMode?true: false">Edit Exam List
    </button>
  </div>
</template>

<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
import moment from 'moment'

export default {
  created:function() {
    this.getExamList();
  },
  components: {
    VueBootstrapDatetimepicker
  },
  name: "examListDetails",
  data: () => ({
    test:'dsqfsd',
    time:null,
    options:{
      format:"YYYY-MM-DD",
      useCurrent: false
    },
    id: 0,
    examDate: null,
    examList: null,
    editMode: false,
    data: null,
    examListUpdated: false,
    errorUpdateExamList: null,
  }),
  methods:{
    getExamList: function(){
      this.id = this.$route.params.id;
      window.backend.Service.GetExamList(this.id).then(data=>{
        this.examList = data;
        this.examDate = this.examList.date_exam;
      });
    },
    removeStudent: function(exam){
      for (var i = 0; i < this.examList.students_exams.length; i++) {
        if(this.examList.students_exams[i].id == exam.id){
          this.examList.students_exams.splice(i, 1);
        }
      }
    },
    editexamList: function(){
      window.backend.Service.UpdateExamList(
        this.examList.id,
        this.examDate = moment(this.examDate).format(),
        this.examList.examiner,
        this.examList.students_exams,
        [],
      ).then(
        data=>{this.data= data},
        err=>{this.errorUpdateExamList = err},
      );
      if(this.errorUpdateExamList != null){
        this.examListUpdated = false;
      }else{
        this.examListUpdated = true;
        this.errorUpdateExamList = null;
        this.switchEditMode();
      }

    },
    switchEditMode: function(){
      this.editMode = !this.editMode;
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