const { expect } = require('chai');

describe('Sample test describe block', function() {
  before(() => {
    console.log('Sample before block executing...')
  });

  it('should test', done => {
    expect(true).to.be.false;
    done();
  })
})
