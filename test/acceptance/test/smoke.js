require('mocha-generators').install();

const Nightmare = require('nightmare');
const expect = require('chai').expect; // jshint ignore:line
const rand = require("randomstring").generate;

const baseUrl = process.env.URL || 'http://localhost:6111';
const bindPw = process.env.PW || 'TestMe123';
const userName = "_" + rand({length: 5, charset: 'alphanumeric'});

describe('smoke test', () => {

  const nightmare = Nightmare();
  const delLink = `a#del_${userName}`;

  it('should login first', function*() {
    yield nightmare
      .goto(baseUrl)
      .wait(() => location.pathname === '/login')
      .type('input#Name', 'Administrator')
      .type('input#Pass', bindPw)
      .click('input#submit')
      .wait(() => location.pathname === '/');
  });

  it('should be able to add user', function*() {
    yield nightmare
      .click('a#add_user')
      .wait('input#Name')
      .type('input#Name', `${userName}`)
      .type('input#DisplayName', `${userName} LastName`)
      .type('input#Password', `${userName}Pass123`)
      .click('input#create')
      .wait(() => location.pathname === '/')
  });

  it('should be able to see just added user', function*() {
    yield nightmare
      .evaluate((link) => {
        let el = document.querySelector(link);
        let nodeList = el.parentNode.parentNode.querySelectorAll('td');
        let res = []
        for (i in nodeList) {
          res.push(nodeList[i].innerHTML)
        }
        return res
      }, delLink)
      .then((tds) => {
        expect(tds[0]).to.equal(userName);
        expect(tds[1]).to.equal(`${userName} LastName`);
        expect(tds[2]).to.equal(' ✓ ')
      });
  });

  it('user delete should work', function*() {
    yield nightmare
      .click(delLink)
      .wait((link) => !document.querySelector(link), delLink)
  });
});
