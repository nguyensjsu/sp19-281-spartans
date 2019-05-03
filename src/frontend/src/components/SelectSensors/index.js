import React, {Component} from 'react';
import Card from '@material-ui/core/Card';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';


const ids = [
    "1234",
    "computer2"
];

const styles = {
    card: {
        padding: "2vh",
        width: "70vw",
        height: "50vh",
        margin: '100px auto '
    }
};

class SelectSensors extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: ids.map(id => ({id: id, checked: false}))
        }
    }

    handleToggle(index) {
        let item = {...this.state.data[index], checked: !this.state.data[index].checked};
        let newData = [...this.state.data.slice(0, index), item, ...this.state.data.slice(index + 1)]
        this.setState({data: newData});
        this.props.setSelectedSensors(newData.filter(item => item.checked).map(item => item.id))

    }

    render() {
        let {data} = this.state;
        return (
            <div>
                <Card style={styles.card}>
                    Refresh Rate :
                    {
                        data.map((id, index) =>
                            <FormControlLabel
                                key={index}
                                control={
                                    <Checkbox
                                        checked={id.checked}
                                        onChange={this.handleToggle.bind(this, index)}
                                    />
                                }
                                label={id.id}
                            />
                        )
                    }
                </Card>
            </div>
        );
    }
}

export default SelectSensors;
