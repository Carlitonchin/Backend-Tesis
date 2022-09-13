import {createStore} from 'vuex'

export default createStore({
    state:{
        user:null
    },
    mutations:{
        SET_USER(state, payload){
            state.user = payload
        }
    }
})