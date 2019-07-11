<template>
  <div id="addStudent">

    <div class="mb-5 mt-3">
      <h1>Add New Student</h1>
    </div>
    <p>{{dataBackDebug}}</p>
    <!--  errors  -->
    <div v-if="error" class="errors">
      <p>"There is an error"</p>
    </div>

    <!-- form -->
    <div class="col-md order-md-1">
      <form class="needs-validation" novalidate="">

        <div class="row">
          <!-- first name -->
          <div class="col-md-6 mb-3">
            <label for="firstName">First name</label>
            <input v-model="firstName" type="text" class="form-control"
            id="firstName" placeholder="" value="" required>
            <div class="invalid-feedback">
              Valid first name is required.
            </div>
          </div>

          <!-- last name -->
          <div class="col-md-6 mb-3">
            <label for="lastName">Last name</label>
            <input v-model="lastName" type="text" class="form-control"
            id="lastName" placeholder="" value="" required>
            <div class="invalid-feedback">
              Valid last name is required.
            </div>
          </div>
        </div>

      <div class="row">

        <!-- maiden name -->
        <div class="col mb-3">
          <label for="address">Maiden Name</label>
          <input v-model="maidenName" type="text" class="form-control" id="maidenName" placeholder="" required>
        </div>

        <!-- phone number -->
        <div class="col mb-3">
          <label for="address">Phone Number</label>
          <input v-model="phoneNumber" type="text" class="form-control" id="maidenName" placeholder="" required>
        </div>
      </div>

      <div class="row">

        <!-- birthday -->
        <datetime placeholder="Exam date"  v-model="birthday"></datetime>

        <!-- <div class="col-4 mb-3">
          <label for="birthday">Birthday</label>
          <input v-model="birthday" type="text" class="form-control" id="birthday" placeholder="" required="false">
        </div> -->

        <!-- birth city -->
        <div class="col-4 mb-3">
          <label for="birthCity">BirthCity</label>
          <select v-model="birthCity" class="custom-select d-block" id="birthCity" required="">
            <option value="oran">Oran</option>
          </select>
        </div>

        <!-- gender -->
        <div class="col-4 mb-3">
          <label for="gender">Gender</label>
          <select v-model="gender" class="custom-select d-block w-100"
                  id="gender" required="">

            <option value="man">Man</option>
            <option value="woman">Woman</option>
          </select>
          <div class="invalid-feedback">
            Please select a valid country.
          </div>
        </div>

        <!-- fileNumber -->
        <div class="col-4 mb-3">
          <label for="fileNumber">FileNumber</label>
          <input v-model="fileNumber" type="text" class="form-control" id="fileNumber" placeholder="" required="false">
        </div>

        <!-- job -->
        <div class="col-4 mb-3">
          <label for="job">Job</label>
          <input v-model="job" type="text" class="form-control" id="job" placeholder="" required="false">
        </div>

        <!-- registred date -->
        <datetime placeholder="Exam date"  v-model="registredDate"></datetime>

        <!-- <div class="col-4 mb-3">
          <label for="registredDate">Registred Date</label>
          <input v-model="registredDate" type="text" class="form-control" id="registredDate" placeholder="" required="false">
        </div> -->

      </div>

      <div class="row">
        <!-- address -->
        <div class="col-12 mb-3">
          <label for="address">Address</label>
          <input v-model="addressStreet" type="text" class="form-control" id="address" placeholder="1234 Main St" required="">
          <div class="invalid-feedback">
            Please enter your shipping address.
          </div>
        </div>

        <!-- department -->
        <div class="col-4 mb-3">
          <label for="department">Department</label>
          <input v-model="department" type="text" class="form-control" id="address" placeholder="Bir el djir" required="">

        </div>

        <!-- city -->
        <div class="col-4 mb-3">
          <label for="city">City</label>
          <select v-model="city" class="custom-select d-block" id="city" required="">
            <option value="oran">Oran</option>
          </select>
        </div>
      </div>

      <!-- button add student -->
      <button @click.prevent="addStudent()" class="btn btn-primary btn-lg btn-block" type="submit">Add Student</button>
    </form>
  </div>
</div>
</template>


<script>
export default {
  name: "addStudent",
  data: () => ({
    error: false,
    fileNumber: "",
    firstName: "",
    lastName: "",
    maidenName: "",
    phoneNumber: "",
    birthday: "",
    birthCity: " ",
    job: "",
    gender: "male",
    country: "algeria",
    city: " ",
    department: "",
    addressStreet: "",
    registredDate: "",
    created: null,
    student:{},
    dataBackDebug: null
  }),
  methods:{
    addStudent:function(){
      if (this.firstName == "" ||
        this.lastName == "" ||
        this.maidenName == "" ||
        this.phoneNumber == "" ||
        this.birthday == "" ||
        this.birthCity == "" ||
        this.job == "" ||
        this.fileNumber == "" ||
        this.gender == "" ||
        this.country == "" ||
        this.city == "" ||
        this.department == "" ||
        this.addressStreet == "" ||
        this.registredDate == ""
      ){
        this.error = true;
        this.dataBackDebug = this.student = {
          "first_name": this.firstName,
          "last_name": this.lastName,
          "maiden_name": this.maidenName,
          "phone_number": this.phoneNumber,
          "birthday": this.birthday,
          "birth_city": this.birthCity,
          "job": this.job,
          "city": this.city,
          "file_number": this.fileNumber,
          "gender": this.gender,
          "country": this.country,
          "department": this.department,
          "address_street": this.addressStreet,
          "registred_date": this.registredDate,
        }
      }else{
        this.error = false;
        this.student = {
          "first_name": this.firstName,
          "last_name": this.lastName,
          "maiden_name": this.maidenName,
          "phone_number": this.phoneNumber,
          "birthday": this.birthday,
          "birth_city": this.birthCity,
          "job": this.job,
          "city": this.city,
          "file_number": this.fileNumber,
          "gender": this.gender,
          "country": this.country,
          "department": this.department,
          "address_street": this.addressStreet,
          "registred_date": this.registredDate,
        }

        window.backend.Service.CreateStudent(this.student).then(data=>{
          this.dataBackDebug= data;
        });
      }
    }
  }
}
</script>
<style scoped>
datetime{
  border: none;
}
</style>