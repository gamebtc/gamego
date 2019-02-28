let rsconf = {
  _id: "gameRs",
  members: [
    {
     _id: 0,
     host: "127.0.0.1:27088"
    },
    {
     _id: 1,
     host: "127.0.0.1:27089"
    },
    {
     _id: 2,
     host: "127.0.0.1:27090"
    }
   ]
}
rs.initiate( rsconf )
rs.conf()
rs.status() 