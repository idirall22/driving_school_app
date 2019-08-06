import moment from 'moment';

export function analyseDataPerMonth(list){
  let listOut = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
  for (var i = 0; i < list.length; i++) {
    let date = moment(list[i]).month()
    switch (date) {
      case 0:
        listOut[0] += 1
        break;
      case 1:
        listOut[1] += 1
        break;
      case 2:
        listOut[2] += 1
        break;
      case 3:
        listOut[3] += 1
        break;
      case 4:
        listOut[4] += 1
        break;
      case 5:
        listOut[5] += 1
        break;
      case 6:
        listOut[6] += 1
        break;
      case 7:
        listOut[7] += 1
        break;
      case 8:
        listOut[8] += 1
        break;
      case 9:
        listOut[9] += 1
        break;
      case 10:
        listOut[10] += 1
        break;
      case 11:
        listOut[11] += 1
        break;
      default:
        break
    }
  }
  return listOut
}