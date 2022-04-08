<template>
  <div>
    <svg class="svgRef"></svg>
  </div>
</template>

<script>
import {onMounted} from "vue";

import axios from "axios";
import * as topojson from "topojson";
import * as d3 from "d3";

export default {
  name: "GooseMap",
  setup() {
    function getGooseData(){
      return axios.get("http://localhost:8081/birds")
        .then(response => {
          return response.data
        })
    }
    
    let height = 800
    let width = 800

    d3.json("http://localhost:8081/state")
      .then(function(state){

        let counties = topojson.feature(state, state.objects.cb_2015_indiana_county_20m)

        let projection = d3.geoConicEqualArea()
          .rotate([86, 0])
          .center([0, 40])
          .parallels([38, 48])
          .scale([7000])
        // let projection = d3.geoAlbersUsa()
        //   .translate([width/2, height/2])
        //   .scale([3000]);          // scale things down so see entire US

        let path = d3.geoPath()
          .projection(projection);

        // Join the FeatureCollection's features array to path elements
        d3.select(".svgRef")
          .append("path")
          .datum(counties)
          .attr("d", path)
          .attr("fill", "none")
          .attr("stroke", "black")
          .attr("stroke-width", "0.5px")

    }).catch(err => {
      throw err;
    })
  
    onMounted(() => {
      let projection = d3.geoConicEqualArea()
        .rotate([86, 0])
        .center([0, 40])
        .parallels([38, 48])
        .scale([7000])

      // pass ref with DOM element to D3, when mounted (DOM available)
      const svg = d3.select(".svgRef");
      svg
        .attr("height", height)
        .attr("width", width)

      getGooseData().then(data => {
        data = data.filter(d => {return !isNaN(Math.sqrt(d.howMany))})

        svg
          .selectAll(".goose") // get all "existing" lines in svg
          .data(data) // sync them with our data
          .join("circle") // create a new "path" for new pieces of data (if needed)
            // everything after .join() is applied to every "new" and "existing" element
            .attr("class", "goose") // attach class (important for updating)
            .attr("cx", function(d) {
              return projection([d.lng, d.lat])[0]
            })
            .attr("cy", function(d) {
              return projection([d.lng, d.lat])[1]
            })
            .attr("r", d => 0.5 * Math.sqrt(Math.PI * d.howMany))
            .attr("fill", "green") // styling
      });
    });

  },

}
</script>
