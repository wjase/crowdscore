import React from 'react';
import {connect} from 'react-redux';
import Form from 'grommet/components/Form';
import Header from 'grommet/components/Header';
import Rating from "react-rating";
import {FormField, Paragraph} from "grommet";

class Vote extends React.Component {
    static propTypes = {}

    render = () => {
        let team = "bob";
        return (
            <div>
                <Form>
                    <Header>{team} Score</Header>
                    <Paragraph>Please provide your input on the followng teams</Paragraph>
                    <FormField label="How much collaboration occurred with this project?">
                        <Rating name="Collaboration" start="1" stop="10" step="2"/>
                    </FormField>
                    <FormField label="How innovative is the product">
                        <Rating name="Innnovation" start="1" stop="10" step="2"/>
                    </FormField>
                    <FormField label="How complete is the product?">
                        <Rating name="Completeness" start="1" stop="10" step="2"/>
                    </FormField>
                    <FormField label="How complete is the product?">
                        <Rating name="Impact" start="1" stop="10" step="2"/>
                    </FormField>
                </Form>
            </div>
        )
    }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({});

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Vote);
