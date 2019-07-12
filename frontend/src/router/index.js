import Vue from 'vue'
import Router from 'vue-router'
import ListStudents from '@/components/ListStudents.vue'
import AddStudents from '@/components/AddStudent.vue'
import AddExamList from '@/components/AddExamList.vue'
import ExamLists from '@/components/ExamLists.vue'
import StudentDetails from '@/components/StudentDetails.vue'
import ExamListDetails from '@/components/ExamListDetails.vue'

import 'bootstrap/dist/css/bootstrap.css';
import 'pc-bootstrap4-datetimepicker/build/css/bootstrap-datetimepicker.css';
import moment from 'moment';


Vue.use(Router)
Vue.filter("capitalize",str => str.charAt(0).toUpperCase() + str.slice(1))
Vue.filter("moment",date => moment(date).format('DD MMMM YYYY'))
Vue.filter("moment2",date => moment(date).format('DD-MM-YYYY'))


export default new Router({
  linkExactActiveClass: 'active',
  routes:[
    {path: "/students/", name: "listStudents",component: ListStudents},
    {path: "/add-student", name: "addStudent",component: AddStudents},
    {path: "/students/:id", name: "studentDetails",component: StudentDetails},

    {path: "/exam-lists", name: "examLists",component: ExamLists},
    {path: "/add-exam-list", name: "addExamList",component: AddExamList},
    {path: "/exam-lists/:id", name: "examListDetails",component: ExamListDetails}
  ]
})