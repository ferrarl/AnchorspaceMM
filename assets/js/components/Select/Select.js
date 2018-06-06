import React from 'react'

const Select = props => {
  var options = props.options.map(option => {
    let display = option[props.displayField]
    return (
      <option key={option.id} value={option.id}>{display}</option>
    )
  })

  return (
    <div className='row justify-content-center'>
      <div className='col-4'>
        <br />
        <select className='custom-select' onChange={props.handleChange} value={props.selectedID}>
          <option value={''}>{props.title}</option>
          {options}
        </select>
        <br />
        <br />
      </div>
    </div>
  )
}

export default Select
