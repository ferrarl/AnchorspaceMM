import { ANCHOR_URL } from '../constants'

export function FetchListings (params) {
  var queryString = ''
  if (params) {
    queryString = params.toString()
  }
  return FetchData(`listings.json?${queryString}`)
}

export function FetchNeighborhoods () {
  return FetchData(`neighborhoods.json`)
}

export function HandleLogout(params, then_cb, error) {
  //TODO: fix this!!!!
  document.cookie = "_dashboard_session=; expires=Thu, 01 Jan 1970 00:00:00 GMT"
  // var queryString = ''
  // if (params) {
  //   queryString = params.toString()
  // }
  // fetch (
  //   `/signout?${queryString}`,
  //   {
  //     method: 'DELETE',
  //   }
  // )
  // .then( res => (then_cb !== undefined ? then_cb(res) : null))
  // .catch( err => (error !== undefined ? error(err) : null))
}

function FetchData (endpoint) {
  return fetch(
    `${ANCHOR_URL}/${endpoint}`,
    { headers: { 'Content-Type': 'application/json' } }
  ).then(response => response.json())
}
