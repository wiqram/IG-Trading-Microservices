module.exports = {
  port: process.env.PORT || 3070,
  db: {
      host: 'mysql',
      database: 'products',
      user: 'root',
      password: 'password',
      port: 3306
  }
};
