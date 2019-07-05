<template>
  <div id="addStudent">
    <div class="mb-5 mt-3">
      <h1>Add New Student</h1>
    </div>
    <div v-if="error" class="errors">
      <p>"There is an error"</p>
    </div>

    <div v-if="created" class="created">
      <p>"student added"</p>
    </div>

    <div class="col-md order-md-1">
      <form class="needs-validation" novalidate="">

        <div class="row">
          <!-- first name -->
          <div class="col-md-6 mb-3">
            <label for="firstName">First name</label>
            <input v-model="first_name" type="text" class="form-control"
            id="firstName" placeholder="" value="" required>
            <div class="invalid-feedback">
              Valid first name is required.
            </div>
          </div>

          <!-- last name -->
          <div class="col-md-6 mb-3">
            <label for="lastName">Last name</label>
            <input v-model="last_name" type="text" class="form-control"
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
          <input v-model="maiden_name" type="text" class="form-control" id="maidenName" placeholder="" required>
        </div>

        <!-- phone number -->
        <div class="col mb-3">
          <label for="address">Phone Number</label>
          <input v-model="phone_number" type="text" class="form-control" id="maidenName" placeholder="" required>
        </div>
      </div>

      <div class="row">

        <!-- birthday -->
        <div class="col mb-3">
          <label for="birthday">Birthday</label>
          <input v-model="birthday" type="text" class="form-control" id="birthday" placeholder="" required="false">
        </div>

        <!-- job -->
        <div class="col mb-3">
          <label for="job">Job</label>
          <input v-model="job" type="text" class="form-control" id="job" placeholder="" required="false">
        </div>

        <!-- gender -->
        <div class="col mb-3">
          <label for="gender">Gender</label>
          <select v-model="gender" class="custom-select d-block w-100"
                  id="gender" required="">

            <option>Man</option>
            <option>Woman</option>
          </select>
          <div class="invalid-feedback">
            Please select a valid country.
          </div>
        </div>

      </div>

      <div class="row">
        <!-- address -->
        <div class="col-md-6 mb-3">
          <label for="address">Address</label>
          <input v-model="address_street" type="text" class="form-control" id="address" placeholder="1234 Main St" required="">
          <div class="invalid-feedback">
            Please enter your shipping address.
          </div>
        </div>

        <!-- city -->
        <div class="col-md-6 mb-3">
          <label for="city">City</label>
          <select class="custom-select d-block" id="city" required="">
            <option value="">Oran</option>
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
    first_name: "",
    last_name: "",
    maiden_name: "",
    phone_number: "",
    birthday: "",
    job: "",
    gender: "",
    address_street: "",
    created: null,
    student:{}
  }),
  methods:{
    addStudent:function(){
      if (this.first_name == "" || this.last_name == "" ||
          this.maiden_name == "" ||this.phone_number == "" ||
          this.birthday == "" || this.job == "" ||
          this.gender == "" || this.address_street == ""){
          this.error = true;
      }else{
        this.error = false;
        this.student = {"first_name": this.first_name, "last_name": this.last_name,
        "maiden_name": this.maiden_name, "phone_number": this.phone_number,
        "birthday": this.birthday, "job":this.job, "gender": this.gender,
        "address_street": this.address_street, "city": "oran"}

        window.backend.Service.CreateStudentMap(this.student).then(err=>{
          if (err == null){
            this.created = true;
          }else{
            this.created = false;
          }
        });
      }
    }
  }
}
</script>
<style scoped>
</style>