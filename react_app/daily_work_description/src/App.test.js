import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'

let container = null

beforeEach(() => {
  container = document.createElement('div')
})

describe('<App />', () => {
  it('リンクを出力する', () => {
    ReactDOM.render(<App />, container)
    expect(container.innerHTML).toContain('Learn React')
  })
})
