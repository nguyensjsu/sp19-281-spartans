import React, {Component} from 'react';
import C3Chart from 'react-c3js';
import 'c3/c3.css';

const data = {
    x: 'x',
    xFormat: '%Y',
    columns: [
        ['x'],
    ],
};

const axis = {
    x: {
        type: 'timeseries',
        localtime: false,
        tick: {
            format: '%H:%M'
        }
    }
};

const styles = {
    margin: "10vh 5vw",
};

class Graph extends Component {
    render() {
        return (
            <div style={styles}>
                <C3Chart data={this.props.data || data} axis={axis}/>
            </div>
        );
    }
}

export default Graph;