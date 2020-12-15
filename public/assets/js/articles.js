const articles = new Vue({
  el: '#article-list',
  data: {
    user: '',
    newArticle: {
      title: '',
      content: ''
    },
    articles: [],
    completedArticles: []
  },
  methods: {
    getArticles() {
      const url = 'api/articles';
      const headers = {'Authorization': `Bearer ${this.getToken()}`};

      fetch(url, {headers}).then(response => {
        if(response.ok) {
          return response.json()
        }
        return []
      }).then(json => {
        this.articles = json
      })
    },
    postArticle() {
      const url = 'api/articles';
      const method = 'POST';
      const headers = {
        'Authorization': `Bearer ${this.getToken()}`,
        'Content-Type': 'application/json; charset=UTF-8'
      };
      const body = JSON.stringify({
        title: this.newArticle.title,
        content: this.newArticle.content,
      });

      fetch(url, {method, headers, body}).then(response => {
        if(response.ok) {
          return response.json()
        }
      }).then(json => {
        if(typeof json === 'undefined') {
          return
        }
        this.articles.push(json);
        this.newArticle = { title: '', content: '' };
      })
    },
    deleteArticle(id) {
      const url = `api/articles/${id}`;
      const method = 'DELETE';
      const headers = {'Authorization': `Bearer ${this.getToken()}`};

      fetch(url, {method, headers}).then(response => {
        if(response.ok) {
          this.articles = this.articles.filter(article => article.id !== id)
        }
      })
    },
    checkArticle(id) {
      const url = `api/articles/${id}/completed`;
      const method = 'PUT';
      const headers = {'Authorization': `Bearer ${this.getToken()}`}
      ;
      fetch(url, {method, headers})
    },
    getToken() {
      return localStorage.getItem('token')
    },
    logout() {
      localStorage.removeItem('token');
      location.href = '/'
    },
  },
  created() {
    const date = new Date();
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.user = claims.name;
    if(claims.exp < Math.floor(date.getTime() / 1000)) {
      this.logout()
    } else {
      this.getArticles()
    }
  },
});
