import React from 'react';
import $ from 'jquery';

class Test2 extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
      </div>
    );
  }
}

Test2.propTypes = {
  foo: React.PropTypes.bool.isRequired,
  bar: React.PropTypes.string.isRequired
}

export default Test2;

// rcg -file=examples/testTwo.js -imports=React:react,$:jquery -mode=class -name=Test2 -props=foo:bool,bar:string
