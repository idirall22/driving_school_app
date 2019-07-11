<template>
  <div id="studentDetails">
    <div class="row mb-5 mt-3">
      <div class="col">
        <h1 class="mb-3">{{student.last_name | capitalize}} {{student.first_name | capitalize}}</h1>
        <button @click="switchEditMode" class="btn btn-warning">Edit mode</button>
      </div>
      <div class="col-3 alert alert-primary pb-0" role="alert">
        <h4>{{student.next_exam | capitalize}}</h4>
        <p v-if="student.last_exam_date != null">Last Exam: {{student.last_exam_date}}</p>
        <p v-if="student.last_exam_date == null"> Not yet</p>
      </div>
    </div>
    <!-- form -->
    <form class="needs-validation" novalidate="">

      <div class="row">

        <!-- first name -->
        <div class="col-6 mb-3">
          <label for="firstName">First name</label>
          <input v-model="student.first_name" type="text" class="form-control"
          id="firstName" placeholder="" value="" required :readonly="!editMode?true: false">
        </div>

        <!-- last name -->
        <div class="col-md-6 mb-3">
          <label for="lastName">Last name</label>
          <input v-model="student.last_name" type="text" class="form-control"
          id="lastName" placeholder="" value="" required :readonly="!editMode?true: false">
        </div>
      </div>

      <div class="row">

        <!-- maiden name -->
        <div class="col mb-3">
          <label for="address">Maiden Name</label>
          <input v-model="student.maiden_name" type="text" class="form-control" id="maidenName"
          placeholder="" required :readonly="!editMode?true: false">
        </div>

        <!-- phone number -->
        <div class="col mb-3">
          <label for="address">Phone Number</label>
          <input v-model="student.phone_number" type="text" class="form-control"
          id="address" placeholder="" required :readonly="!editMode?true: false">
        </div>

      </div>

      <div class="row">

        <!-- birthday -->
        <div class="col-4 mb-3">
          <label for="birthday">Birthday</label>
          <vue-bootstrap-datetimepicker v-model="birthday" :config="options"
          required :disabled="!editMode?true: false" ></vue-bootstrap-datetimepicker>
        </div>

        <!-- birth city -->
        <div class="col-4 mb-3">
          <label for="birthCity">BirthCity</label>
          <select v-model="student.birth_city" class="custom-select d-block" id="birthCity"
          required :disabled="!editMode?true: false">
          <option value="oran">Oran</option>
          </select>
        </div>

        <!-- Country -->
        <div class="col mb-3">
          <label for="country">Country</label>
          <input v-model="student.country" type="text" class="form-control"
          id="country" placeholder="" required :readonly="!editMode?true: false">
        </div>

        <!-- gender -->
        <div class="col-4 mb-3">
          <label for="gender">Gender</label>
          <select v-model="student.gender" class="custom-select d-block w-100"
          id="gender" required :disabled="!editMode?true: false">

          <option value="man">Man</option>
          <option value="woman">Woman</option>
          </select>
        </div>

        <!-- registred date -->
        <div class="col-4 mb-3">
          <label for="registredDate">Registred Date</label>
          <vue-bootstrap-datetimepicker v-model="registredDate" :config="options"
          readonly ></vue-bootstrap-datetimepicker>
        </div>
      </div>
      <div class="row">
        <!-- address -->
        <div class="col-12 mb-3">
          <label for="address">Address</label>
          <input v-model="student.address_street" type="text" class="form-control" id="address" placeholder="1234 Main St"
          required :readonly="!editMode?true: false">
        </div>

        <!-- department -->
        <div class="col-4 mb-3">
          <label for="department">Department</label>
          <input v-model="student.department" type="text" class="form-control" id="address" placeholder="Bir el djir"
          required :readonly="!editMode?true: false">
        </div>

        <!-- city -->
        <div class="col-4 mb-3">
          <label for="city">City</label>
          <select v-model="student.city" class="custom-select d-block" id="city"
          required :disabled="!editMode?true: false">
          <option value="oran">Oran</option>
          </select>
        </div>
      </div>


      <div class="row">
        <!-- fileNumber -->
        <div class="col-4 mb-3">
          <label for="fileNumber">FileNumber</label>
          <input v-model="student.file_number" type="text" class="form-control" id="fileNumber" placeholder="" required
          :readonly="!editMode?true: false">
        </div>
        <!-- job -->
        <div class="col-4 mb-3">
          <label for="job">Job</label>
          <input v-model="student.job" type="text" class="form-control" id="job" placeholder=""
          required :readonly="!editMode?true: false">
        </div>

      </div>
      <!-- button add student -->
      <button @click.prevent="editStudent()"
      class="btn btn-primary btn-lg btn-block"
      type="submit" :disabled="!editMode?true: false">Edit Student</button>
    </form>
  </div>
</template>


<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
import moment from 'moment';

export default {
  name: "studentDetails",
  created: function(){
    this.getStudent();
  },
  components: {
    VueBootstrapDatetimepicker
  },
  data: () => ({
    options:{
      // format:"DD/MM/YYYY",
      useCurrent: false
    },
    editMode:false,
    student:{},
    birthday: null,
    registredDate: null,
    id: null
  }),
  methods:{
    getStudent: function(){
      this.id = this.$route.params.id;
      window.backend.Service.GetStudent(this.id, "").then(data=>{
        this.student = data;
      });
      this.birthday = this.toMoment(this.student.birthday);
      this.registredDate = this.toMoment(this.student.birthday);
    },
    switchEditMode: function(){
      this.editMode = !this.editMode;
    },
    editStudent: function(){
      window.backend.Service.UpdateStudent(this.student);
    },
    toMoment:function(date){
      return moment(date).format('DD MMMM YYYY');
    }
  }
}
</script>

<style scoped>
.contentLeft{
  border-right: 1px solid #333;
}
</style>
