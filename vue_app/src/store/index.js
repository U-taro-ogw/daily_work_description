import Vue from 'vue'
import Vuex from 'vuex'

import moment from 'moment'

Vue.use(Vuex)

const state = {
  monthDays: []
}

const mutations = {
  setMonthDays(state, today) {
    let endOfMonth = moment(today).endOf('month').get('date')
    state.monthDays = Array(endOfMonth).fill().map((_, k) => k + 1)
  }
}

const actions = {
  setMonthDays(context) {
    context.commit('setMonthDays', moment())
  }
}

export default new Vuex.Store({
  state,
  mutations,
  actions,
  modules: {
  }
})
