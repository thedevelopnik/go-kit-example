Vue.use(Vuetify)

new Vue({
  el: '#app',
  data() {
    return {
      newTodo: '',
      todos: []
    }
  },
  methods: {
    submit: function() {
      this.todos.push(this.newTodo)
      this.newTodo = ''
    },
    remove: function(index) {
      this.todos.splice(index, (index + 1))
    },
    getTodos: function() {
      axios.get('http://localhost:3030/gettodos')
      .then(res => {
        res.data.v.forEach(todo => {
          this.todos.push(todo.name)
        })
      })
    }
  }
})
