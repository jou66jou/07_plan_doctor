var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

router.get('/diseaseMap', function(req, res, next) {
  res.render('gmap', { title: '病害地圖' });
});

router.get('/news', function(req, res, next) {
  res.render('main_news', { title: '病害地圖' });
});

router.get('/community', function(req, res, next) {
  res.render('main_community', { title: '病害地圖' });
});

router.get('/bulletin', function(req, res, next) {
  res.render('main_bulletin', { title: '病害地圖' });
});

module.exports = router;
