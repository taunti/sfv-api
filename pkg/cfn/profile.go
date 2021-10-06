package cfn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var Leagues = map[string]string{
	"0":  "Rookie",
	"1":  "Bronze",
	"2":  "Super Bronze",
	"3":  "Ultra Bronze",
	"4":  "Silver",
	"5":  "Super Silver",
	"6":  "Ultra Silver",
	"7":  "Gold",
	"8":  "Super Gold",
	"9":  "Ultra Gold",
	"10": "Platinum",
	"11": "Super Platinum",
	"12": "Ultra Platinum",
	"13": "Diamond",
	"14": "Super Diamond",
	"15": "Ultra Diamond",
	"16": "Master",
	"17": "Grand Master",
	"18": "Ultimate Grand Master",
	"19": "Warlord",
}

type Profile struct {
	doc *html.Node
}

func (p Profile) getContentAsString(xpath string) string {
	element := htmlquery.FindOne(p.doc, xpath)
	if element == nil {
		return ""
	}

	return strings.TrimSpace(htmlquery.InnerText(element))
}

func (p Profile) getContentAsNumber(xpath string) int {
	element := htmlquery.FindOne(p.doc, xpath)
	content := htmlquery.InnerText(element)

	i, err := strconv.Atoi(content)
	if err != nil {
		fmt.Println(err)
	}

	return i
}

func (p Profile) GetDevice() string {
	const xpath = `//div[contains(@class, 'playerStatus')]/div[1]/p`
	return p.getContentAsString(xpath)
}

func (p Profile) GetStatus() string {
	const xpath = `//div[contains(@class, 'playerStatus')]/div[2]/p`
	return p.getContentAsString(xpath)
}

func (p Profile) GetLeague() string {
	const xpath = `//div[contains(@class, 'leagueInfo')]/dl[1]/dd/img`

	element := htmlquery.FindOne(p.doc, xpath)
	content := htmlquery.SelectAttr(element, "src")

	regex := regexp.MustCompile(`.*\/(?P<league>\d+)\.png.*`)
	value := regex.FindStringSubmatch(content)[1]

	return Leagues[value]
}

func (p Profile) GetRegion() string {
	const xpath = `//p[contains(@class, 'rating')]/img`

	element := htmlquery.FindOne(p.doc, xpath)
	content := htmlquery.SelectAttr(element, "src")

	regex := regexp.MustCompile(`.*\/flags\/(?P<flag>[^.]+)\.png.*`)
	value := regex.FindStringSubmatch(content)[1]

	return value
}

func (p Profile) GetFighterId() string {
	const xpath = `//div[contains(@class, 'fighterInfo')]/dl/dd`
	return p.getContentAsString(xpath)
}

func (p Profile) GetTwitter() string {
	const xpath = `//a[contains(@class, 'twitterLink')]`
	return strings.Replace(p.getContentAsString(xpath), " ", "", 1)
}

func (p Profile) GetTitle() string {
	const xpath = `//div[contains(@class, 'fighterInfo')]/p[2]/span`
	return p.getContentAsString(xpath)
}

func (p Profile) GetDojo() string {
	const xpath = `//div[contains(@class, 'fighterInfo')]/dl[2]/dd`
	return p.getContentAsString(xpath)
}

func (p Profile) GetCharacterName() string {
	const xpath = `//div[contains(@class, 'characterInfo')]/dl[1]/dd/a/text()`
	return p.getContentAsString(xpath)
}

func (p Profile) GetCharacterLevel() int {
	const xpath = `//div[contains(@class, 'characterLevel')]/dl/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetPlayerLevel() int {
	const xpath = `//div[contains(@class, 'playerInfo')]/dl[1]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetRanking() int {
	const xpath = `//div[contains(@class, 'playerInfo')]/dl[2]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetMostRecentMatch() string {
	const xpath = `//div[contains(@class, 'playData')]/dl[1]/dd`
	return p.getContentAsString(xpath)
}

func (p Profile) GetFirstMatch() string {
	const xpath = `//div[contains(@class, 'playData')]/dl[2]/dd`
	return p.getContentAsString(xpath)
}

func (p Profile) GetAccountAge() string {
	first := ParseDate(p.GetFirstMatch())
	last := ParseDate(p.GetMostRecentMatch())

	days := DiffInDays(first, last)

	return fmt.Sprintf(
		"%.f days (%dy, %d months)",
		days,
		int(days/365),
		(int(days)%365)/30,
	)
}

func (p Profile) GetLeaguePoints() int {
	const xpath = `//div[contains(@class, 'leagueInfo')]/dl[2]/dd`
	lp := p.getContentAsString(xpath)

	regex := regexp.MustCompile(`(?P<points>\d+).*`)
	value := regex.FindStringSubmatch(lp)[1]

	i, _ := strconv.Atoi(value)
	return i
}

func (p Profile) GetRankedMatchesWins() int {
	const xpath = `//li[contains(@class, 'battleType rank')]/div[contains(@class, 'battleNumber')]/dl[2]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetRankedMatchesLosses() int {
	const xpath = `//li[contains(@class, 'battleType rank')]/div[contains(@class, 'battleNumber')]/dl[3]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetRankedMatchesTotal() int {
	const xpath = `//li[contains(@class, 'battleType rank')]/div[contains(@class, 'battleNumber')]/dl[1]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetRankedMatchesRatio() float32 {
	wins := p.GetRankedMatchesWins()
	total := p.GetRankedMatchesTotal()

	if wins+total == 0 {
		return 0
	}
	return 100 * (float32(wins) / float32(total))
}

func (p Profile) GetCasualMatchesWins() int {
	const xpath = `//li[contains(@class, 'battleType casual')]/div[contains(@class, 'battleNumber')]/dl[2]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetCasualMatchesLosses() int {
	const xpath = `//li[contains(@class, 'battleType casual')]/div[contains(@class, 'battleNumber')]/dl[3]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetCasualMatchesTotal() int {
	const xpath = `//li[contains(@class, 'battleType casual')]/div[contains(@class, 'battleNumber')]/dl[1]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetCasualMatchesRatio() float32 {
	wins := p.GetCasualMatchesWins()
	total := p.GetCasualMatchesTotal()

	if wins+total == 0 {
		return 0
	}
	return 100 * (float32(wins) / float32(total))

}

func (p Profile) GetBattleLoungeMatchesWins() int {
	const xpath = `//li[contains(@class, 'battleType lounge')]/div[contains(@class, 'battleNumber')]/dl[2]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetBattleLoungeMatchesLosses() int {
	const xpath = `//li[contains(@class, 'battleType lounge')]/div[contains(@class, 'battleNumber')]/dl[3]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetBattleLoungeMatchesTotal() int {
	const xpath = `//li[contains(@class, 'battleType lounge')]/div[contains(@class, 'battleNumber')]/dl[1]/dd`
	return p.getContentAsNumber(xpath)
}

func (p Profile) GetBattleLoungeMatchesRatio() float32 {
	wins := p.GetBattleLoungeMatchesWins()
	total := p.GetBattleLoungeMatchesTotal()

	if wins+total == 0 {
		return 0
	}
	return 100 * (float32(wins) / float32(total))
}

func (p Profile) GetTotalMatches() int {
	return (p.GetRankedMatchesTotal() + p.GetCasualMatchesTotal() + p.GetBattleLoungeMatchesTotal())
}
