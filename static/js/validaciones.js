async function GraficoUltimosSeisMeses() {

    const response = await fetch("/articulo/graficoUltimosSeisMeses");
	const data = await response.json();

    const mesNombres = data.map(item => item.MesNombre);
    const rentabilidad = data.map(item => item.Rentabilidad);
    const montoMayorista = data.map(item => item.MontoMayorista);
    const montoEcommerce = data.map(item => item.MontoEcommerce);
  
    Highcharts.chart('idGraficoUltimosSeisMeses', {
        chart: {
            type: 'column'
        },
        title: {
            text: 'Ventas (mayorista y ecommerce) y Rentabilidad de los Ultimos 6 Meses.',
            align: 'left',
            style: {
                color: 'black',
                fontSize: '25px'
            }
        },
        xAxis: {
            categories: mesNombres,
            labels: {
                style: {
                    color: 'black',
                    fontSize: '20px'
                }
            }
        },
        yAxis:[{
            min: 0,
            title: {
                text: 'Ventas',
                style: {
                    color: 'black',
                    fontSize: '14px'
                }
                
            },
            stackLabels: {
                enabled: true,
                style: {
                    fontWeight: 'bold',
                    color: ( // theme
                        Highcharts.defaultOptions.title.style &&
                        Highcharts.defaultOptions.title.style.color
                    ) || 'gray',
                    textOutline: 'none'
                }
            }
        }, {
            title: {
                text: 'Rentabilidad',
                style: {
                    color: 'black',
                    fontSize: '14px'
                }
            },
            labels: {
                format: '{value}%',
            },
            opposite: true
        }],
        legend: {
            align: 'left',
            x: 70,
            verticalAlign: 'top',
            y: 70,
            floating: true,
            backgroundColor:
                Highcharts.defaultOptions.legend.backgroundColor || 'white',
            borderColor: '#CCC',
            borderWidth: 1,
            shadow: false,
            itemStyle: {
                color: 'black',
                fontSize: '18px'
            }
        },
        tooltip: {
            headerFormat: '<b>{point.x}</b><br/>',
            pointFormat: '{series.name}: {point.y}<br/>Total: {point.stackTotal}'
        },
        plotOptions: {
            column: {
                stacking: 'normal',
                dataLabels: {
                    enabled: true
                }
            }
        },
        series: [{
            name: 'Mayorista',
            data: montoMayorista,
            color: '#00adcb',
            dataLabels: {
                enabled: true,
                color: 'black',
                fontSize: '20px', // Aumenta el tamaño de las etiquetas de datos para la serie "Ecommerce"
                textOutline: 'none'
            }
        }, {
            name: 'Ecommerce',
            data: montoEcommerce,
            color: '#292d67',
            dataLabels: {
                enabled: true,
                color: 'black',
                fontSize: '20px', // Aumenta el tamaño de las etiquetas de datos para la serie "Ecommerce"
                textOutline: 'none'
            }
        }, 
        {
            type: 'spline',
            name: 'Rentabilidad',
            data: rentabilidad,
            color: '#ca0372',
            lineColor: '#ca0372',
            yAxis: 1,
            marker: {
                enabled: true, // Activa los marcadores personalizados
                radius: 4, // Aumenta el tamaño de los marcadores
                lineWidth: 2,
                lineColor: '#ca0372',
                fillColor: '#fff'
            },
            tooltip: {
                pointFormat: '{series.name}: {point.y}%',
            }
        }]
    });

  }

  async function GraficoVentasActualMesAnterior(){


    Highcharts.chart('idGrafico', {
        title: {
            text: 'Combination chart'
        },
        xAxis: {
            categories: ['Apples', 'Oranges', 'Pears', 'Bananas', 'Kiwi'],
        },
        colors: ['#1de9b6', '#1dc4e9', '#A389D4'],
        labels: {
            items: [{
                html: 'Total fruit consumption',
                style: {
                    left: '50px',
                    top: '18px',
                    color: (Highcharts.theme && Highcharts.theme.textColor) || 'black'
                }
            }]
        },
        series: [{
            type: 'column',
            name: 'Jane',
            data: [3, 2, 1, 3, 4]
        }, {
            type: 'column',
            name: 'John',
            data: [2, 3, 5, 7, 6]
        }, {
            type: 'column',
            name: 'Joe',
            data: [4, 3, 3, 9, 0]
        }, {
            type: 'spline',
            name: 'Average',
            data: [3, 2.67, 3, 6.33, 3.33],
            color: '#f44236',
            lineColor: '#f44236',
            marker: {
                lineWidth: 2,
                lineColor: '#f44236',
                fillColor: '#fff'
            }
        }, {
            type: 'pie',
            name: 'Total consumption',
            data: [{
                name: 'Jane',
                y: 13,
                color: '#1de9b6'
            }, {
                name: 'John',
                y: 23,
                color: '#1dc4e9',
            }, {
                name: 'Joe',
                y: 19,
                color: '#A389D4',
            }],
            center: [100, 80],
            size: 100,
            showInLegend: false,
            dataLabels: {
                enabled: false
            }
        }]
    });
                     
  }

async function GraficoTortaEstado(){
    

    const response = await fetch("/articulo/graficoStockEstado");
	const data = await response.json();

    const transformedData = data.map((item) => {
      return {
        name: item.Estado,
        y: parseFloat(item.PorcentajeTotal),
      };
    });

// Build the chart
Highcharts.chart('idGraficoTortaEstado', {
    chart: {
        plotBackgroundColor: null,
        plotBorderWidth: null,
        plotShadow: false,
        type: 'pie'
    },
    title: {
        text: 'Stock de Artículo por Estado',
        align: 'left',
        style: {
            color: 'black',
            fontSize: '25px'
        }
    },
    tooltip: {
        pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
    },
    accessibility: {
        point: {
            valueSuffix: '%'
        }
    },
    plotOptions: {
        pie: {
          allowPointSelect: true,
          cursor: 'pointer',
          dataLabels: {
            enabled: true,
            format: '<b>{point.name}</b>: {point.percentage:.1f} %',
            connectorColor: 'silver',
            style: {
              fontSize: '16px',
              color: 'black'
            }
          }
        }
      },
    series: [{
        name: 'Estados',
        data: transformedData
    }]
});


} 

async function GraficoTortaCategoria(){

    const response = await fetch("/articulo/graficoStockCategoria");
	const data = await response.json();

    const transformedData = data.map((item) => {
      return {
        name: item.DescripcionCategoria,
        y: parseFloat(item.PorcentajeTotal),
      };
    });

    
    // Build the chart
    Highcharts.chart('idGraficoTortaCategoria', {
        chart: {
            plotBackgroundColor: null,
            plotBorderWidth: null,
            plotShadow: false,
            type: 'pie'
        },
        title: {
            text: 'Stock de Artículo por Categoría',
            align: 'left',
            style: {
                color: 'black',
                fontSize: '25px'
            }
        },
        tooltip: {
            pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
        },
        accessibility: {
            point: {
                valueSuffix: '%'
            }
        },
        plotOptions: {
            pie: {
                allowPointSelect: true,
                cursor: 'pointer',
                dataLabels: {
                    enabled: true,
                    format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                    connectorColor: 'silver'
                }
            }
        },
        series: [{
            name: 'Share',
            data:transformedData
        }]
    });

} 

//Para llamar a los demas funciones
async function init() {
    GraficoUltimosSeisMeses();
    GraficoVentasActualMesAnterior();
    GraficoTortaEstado();
    GraficoTortaCategoria();
}

  document.addEventListener('DOMContentLoaded', init);