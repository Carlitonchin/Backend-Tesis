export const API_URL = "http://172.18.0.5:8080/"
export const STUDENT_ROLE = "Estudiante"
export const WORKER_ROLE = "Clasificador"
export const ADMIN_ROLE = "Administrador"
export const LEVEL1_SPECIALIST = "Especialista Nivel 1"
export const LEVEL2_SPECIALIST = "Especialista Nivel 2"

export const UNAUTH_STATUS = 401

export const HEADERS = {
    'Content-Type': 'application/json'
  }

export const QuestionsStatusDict = {
  1:'Preguntas sin clasificar',
  2:'Preguntas clasificadas (Nivel 1)',
  3:'Preguntas clasificadas (Nivel 2)',
  4:'Preguntas clasificadas (Nivel Admin)',
  5:'Preguntas resueltas'
}