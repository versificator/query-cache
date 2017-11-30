package main

import (
	"testing"
	"fmt"
)

func TestParse_1(t *testing.T) {
	rawJSON := `{"$query" : "SELECT cutToFirstSignificantSubdomain(concat('http://', domain)) AS domain_top, SUM(ad_request) AS ad_requests, SUM(ad_opportunity) AS ad_opportunities FROM stat.impression PREWHERE ((date >= '$startDay') AND (date <= '$endDay')) AND (cmp_id IN ((SELECT toUInt32(id) FROM dictionary.campaign WHERE is_vertamedia = 1 GROUP BY id) AS _subquery1)) WHERE (environment IN 0) AND ((TIMESTAMP >= $startHour) AND (TIMESTAMP <= $endHour)) GROUP BY domain_top HAVING (ad_requests != 0) OR (ad_opportunities != 0)d",
      "$startDay":"2017-09-01",
      "$endDay":"2017-09-30",
      "$startHour":"1504224000",
      "$endHour":"1506812400",
      "$groupBy": ["column1",
                    "column2"]
}`

	result := parse([]byte(rawJSON))
	fmt.Println(rawJSON)
	fmt.Println(toString(result))
	_ = result
}

func TestParse_2(t *testing.T) {
	rawJSON := `{
   "$query":"select * from default.stats2 where Date >= '$startDay' and Date <= '$endDay' ",
   "$startDay":"2017-04-01",
   "$endDay":"2017-04-02",
   "$startHour":"",
   "$endHour":""
}`

	result := parse([]byte(rawJSON))
	fmt.Println(rawJSON)
	fmt.Println(toString(result))
	_ = result
}