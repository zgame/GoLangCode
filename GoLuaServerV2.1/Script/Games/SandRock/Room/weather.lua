-----------------------------------------------
--- 天气
--    -1    全天气    All
--    1	晴天	SunShine
--    2	雨天	Rain
--    3	沙尘暴	SandStorm
--    4	阴天	OvercastSky
--    5	炎热天	HotDay
-----------------------------------------------


function SandRockRoom:GetWeather()
    return self.weather
end

function SandRockRoom:UpdateWeather()
    local ran = ZRandom.GetRandom()
end