import React from 'react';
import { connect } from 'react-redux';
import Form from 'grommet/components/Form';
import Form from 'grommet/components/Header';
import Fields from 'grommet/components/Fields';
import NumberInput from 'grommet/components/NumberInput';


class Vote extends React.Component {
  static propTypes = {

  }
  render = () => (
    <div>
      <Form>
        <Header>{team} Score</Header>
        <Paragraph>Please provide your input on the followng teams</Paragraph>
        <Fields>
          <NumberInput name="Collaboration" min="1" max="10" />
          <NumberInput name="Innnovation" min="1" max="10" />
          <NumberInput name="Completeness" min="1" max="10" />
          <NumberInput name="Impact" min="1" max="10" />
        </Fields>
        </Form>
    </div>
  )
}

const mapStateToProps = state => ({

});

const mapDispatchToProps = dispatch => ({

});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Vote);
