import React from 'react';
import $ from 'jquery';

const Test = ({ foo, bar }) => {
  return (
    <div>
    </div>
  );
}

Test.propTypes = {
  foo: React.PropTypes.bool.isRequired,
  bar: React.PropTypes.string.isRequired
}

export default Test;

// rcg -file=examples/test.js -imports=React:react,$:jquery -mode=func -name=Test -props=foo:bool,bar:string
