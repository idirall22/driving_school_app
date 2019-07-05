import Vue from 'vue'
import Router from 'vue-router'
import ListStudents from '@/components/ListStudents.vue'
import AddStudents from '@/components/AddStudent.vue'
import AddExamList from '@/components/AddExamList.vue'
import StudentDetails from '@/components/StudentDetails.vue'

Vue.use(Router)

export default new Router({
  routes:[
    {path: "/", name: "listStudents",component: ListStudents},
    {path: "/add-student", name: "addStudent",component: AddStudents},
    {path: "/add-exam-list", name: "addExamList",component: AddExamList},
    {path: "/:id", name: "studentDetails",component: StudentDetails}
  ]
})