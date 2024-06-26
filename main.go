package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/guregu/null/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	illegal_hash = []string{
		"008c13e1e7614f0e5a7fd0894d11f1e370b1e868179399dfe94e50849e88fbc4",
		"0481b2e5006d4a9036d23f0f68b730f646f11c556596ed694bfbf30b2157715b",
		"051a1f7bb5fbd30f63b14c70d1f4587b07ff2840159e5c2ac8aef74ee80b798b",
		"088879daf9e8d227432a09675db6691fd3dbe9bc9f09faa1db16c0e9a407342f",
		"0e86ecb1edb1d4f9d2038eb7213a1ac71ef06af5d16dea87a250bb50a4c8ff4f",
		"0f563fdb40e76614708545b7bc0dcdc344a013347d0d62827233f016de5e9696",
		"109ca1f619c85858720b5b86db1aaf57915e75442cd2ed915580b087ce318187",
		"1255914507d9d4575511283ca16f946585978b0abe82cbc7cc82ec58f6b4ffb1",
		"18977b4c53df395b4d52c2bf9e9582d8c08c103c302b0a663eb54d35a1b62075",
		"195b761f833ea658e682fc739faf09ef786e510863a76c7ef63afaee6561caee",
		"1cb555053cd6574d72c6eb12680b466d70655d5d6906ab19eea3c5451ac42afd",
		"1d1718da778c17836b17a17f4e69a9bb0e6c1f1dbedb9364dd5ef3fa1fc0e3f2",
		"21615dfba2b890b2f3852bad9d9920ea9295ffb59b67a96f6a411c37532ca1ab",
		"25371531e09f38519aabd8d0013fdb6ef46eaff504089ee7a954f4ecd059c364",
		"2a7c473803c1eeb17234de003fc7fd4679d1d595ee97541203f6ee66c5d3f17a",
		"30f9ec8af4d4e75588d6211d3c5358b22b313fbcfcc068f954635fbb7d874902",
		"33e8de8561fada481619273ab4f8e07f3b6e6914719120945f887cb117896a4d",
		"34891d955cab968991b368d56cb5b6ba90dbb50b6d7cbf0f7fdfc1b8a0746875",
		"36d2f49eca6a29570e42514c80bf9f0aa2cf0408d23e62fcdd0b9db7471d1bd2",
		"3ba1cdff2d592da3dcf5b898806244f5325d099e30346949711b4bf87978d7fa",
		"3c7611fc34876e00435730a7e136c78339cfc62b68ce7b13fc8254b3b0a99ed4",
		"40e65c9f1e520a2518fcba4a1ee0a24007cf8acb299cce4caed79facb1865083",
		"42703527ecf9176b5867b923ebc9dc90d55f63758b698721484d9e103433b349",
		"44ecace7bc1bba29728a1a7f360e9b96ad2a0ade3a0a177f96d80342ded62166",
		"45806fc75d45437b4087a727b57f94161a8fdbcd7100e75caa7ea458e676443a",
		"4651fbb127fc7110638f7faa38aaa823ece0f1b36c06aa4058e757fa5661d829",
		"4662ff99e0b0c8d28275ab31d712f67c2956baba323ee9810fc515af8dfb5319",
		"4684f890684bb0a5ced8442272a7476a6256fd9b57151e5e94a58bfa74825c90",
		"4885118068dbe08fa59ca19ea91196cb4faed922ef17de4befed504b2a8a62de",
		"503f98936e5ef9918bf5939447d7126e3dc003c1d77004453e62663a97473d74",
		"5267975bb6630b18dbd4b0c77aa5b7a2e54c36422c5cfd569ba3605d83950996",
		"543d033f57c45cf04927a1fa6f7a57818231afb8c65949deee2b9044eca1b82e",
		"57615f3605409a0c59337a695f2e7aedef9c5b51904f5409880bf78dea049f24",
		"5a0374b49dadb42052239cfa6b675e7cdd84f8def088d5bd7f9b2d608d419779",
		"6416bc12580d60da78b2ee0a8437c3636bfc8158ef6c573171a51c3f34c07567",
		"652b68e4cb7a2629029f350cb41ec713eb60827b37c02293bdb09ff696a423d2",
		"68e0605e4d1660686da40f27efc1a39ebdac945803e05d44383b42a985cae1ba",
		"6a0e923edd2f3fa989c69dfa29b3168da433156c4dfd9cc5c6b81d42da597c97",
		"6acc85ed6448de152ca2830cf484e877a5b767675695aa1a95cf0ab2527f3c81",
		"6bc205746f49450d9dbc9ba4fad2959a3bb38dc898afac317d4e2922baaf9bf8",
		"6c3e888ac7f3e460148fac3d41f16d4ce2ce2893a85ef6fdfa5e6481d43ee80f",
		"6d28469defa4f54c03c3955eb191534b4668ac9544c2240436f1e598ccd16f87",
		"6edc047567e940a776cfeccb960a5ad9bb6e685ceab56be532046ff510830bc8",
		"6ff2b347c6c37a492832be3b5a737ec4b761d5f5264f6edb8fc68bf6bba14ea5",
		"703b9436e834901a9dd5671e8917c71c1a3d67a57b9464e411aa5e01f054cabc",
		"7ae0b585b29f78ca3a7d9d3720e3c87e10745414dd101b40da8ee628d8086f5d",
		"7cd13ff275af2b6b40ab57e8c16cf870aec308e8c299ac3104becfdcc45165fb",
		"7ec33424bb7686bd7616a8b58596a43b361358bfebd79cfdfd5b2df2b7bec81a",
		"7efccbf9bcef2b4922ec1827064f18b9941b979b2f4cf768bdf5f2da6b0d5742",
		"84045f63acae44b540c876b9dafc0e7c5c273f47dc7ef1dad5a00526051e8253",
		"84b7aff662033347fc578b21f3353bfc3731c0f6b912357b5ad5c3bbf84e76ba",
		"89850b05754f72e714e7148c795d1c87cdd68955563b80db7758256d687ee499",
		"8991fa2e9bb8b67c5a66c7e808289a6ad9405558687195b008f84656c7f1a1f7",
		"89ddf7f9b04aa96402340184e2a5e5dfe8c611626c53f2809ac599b61f8084a1",
		"8a346dc2ee417105775e8f1bbba80368e509bf650bf7dff8ad689e29a9bc8e5c",
		"8ca8e9b68e9cfc90821631fe4b8ab9500d134b0a77a5e609e6460d57feebf084",
		"90cae6753c2e52c3a44b9b1c55f70a9c45698a60a7d990ef2144eaf4a1215c92",
		"9461f821c3e4593a5baff295f813447413587ab84f94c631ef031ce37f0c4cab",
		"9489526e53064330adf9d44f70534aa00b42dfbc0e15a5105928916ca4e4fc12",
		"9997f63c53fb7315eb53359847a68d2171e4d09303b5ab128903a617b7b257df",
		"9a73ac01dd01ec509937c3c29ad7246be243cbb020bcc992e64667253b7a3ea8",
		"a0d3ec3208d165d1dfd622dfa9a42f373038505cbf8217ee150f86af1028d194",
		"a2106165e905d4680515b74ce9e9bbac57869cd8660b03333d3498096af415f1",
		"a23a2642f69ecc40f3bbd56fb0c3856b3b4227b25e497c678545a304b5bbaeab",
		"a2dbf18796e652a8e46f6debe95e0127cf9fa6d2fc43f7d0a4f7c541f624a54d",
		"a38f398b6c0883b7409b9c632e52ca5f4df18f97bc86bdcdc84c8b4733ed565a",
		"a3916f1f688a49a9898f6e18bf63e0c0725c1eb49d61ba1ccd644473f5747fbd",
		"a5bdf8c65795ae982e225f3f7684d5813a3831eeeea105cdfcc897bc9735da78",
		"a735d855e6af27d679c4482a32b4475e3625b49ee174d84bfddd85ccea554aba",
		"a7d1c3a1d9fea2097332bcbe30f500d9250b7f50b2c56ae318feb510842e2c9f",
		"a927d10371c5325425ba7c856a7afa31c38a652ff89013986b392b8aaf5b5f0a",
		"a987011f7a7ccf470bdbdd322e8afeb149712ee0088150704285ac49669324a8",
		"a9d7d0afa185d6ff05e87445bbed95aeb2c0d449fe828100daf57e69b1c98a2f",
		"ac69f545a16e1fa65f81de4b17a53b8faa5d1c3756f3207a0c44eb39e0e4cc1d",
		"ad4909acd3d13a2a2e2e369842e464751f5f52e5073cb74f87b39cedcdc62c2c",
		"ae7fe5a807bedb30eeea17d1bbaf07cdc06115f06d24430f3a8542b94d046535",
		"b4e9a04dbd6c0471382ece0775c15d84d4c0b730f948dd5f7c10714e3b76ff0e",
		"b58bbd88bff4dd1f03cddf09a8417d08e4e29acfea47c93748b00f8069fea4d1",
		"b745eaf576827ec502b0ab845c8a48c7592dec3b345019d61f9daa3eeab3e01b",
		"b7bfffa677527eb5090229f3e807a05c7a7c46c1fd8143dbce9ba42e1d5ab234",
		"b92211b6bc29c4ba6e7a44608839ed0bc31d210e198d4d16c413de290a646130",
		"b94bd79bc199d0dd071eca0bb5f1236f48d566e6f27120114af75fa885e490b2",
		"bbf85a60c50850f7ddbe714f06676b0414d25740fa5d9bb5a9e6109e41a0d27e",
		"c0c5e4a448adc9f851d4f3b8c8ff884e6157c4b1c0eb0474f2ab3fcd6ce22299",
		"cc3110b48214912c63c9b58e7333b9e3d9d0dbe3f030f35be30c7d532dca5420",
		"d245d56c5fb2f5ba245d9dec7deea45af0199a51691eafdd5eaa136054c43b9c",
		"d2d1bf0b35db942fb5656dee739d99c7d190d8a1d39f4b4a0084237f833ec896",
		"d5238ace28a4594425213ed34471b87dd0f8c0507e533af34ee0188f615657a3",
		"d561f2b7643d529b7b7916a87ee099b1679d56f3b24bb98e75ba6117c3d7ef6c",
		"d6bb15b687d285a9158ae7307aab559a95e6a3ba1919320cf05557d2395508e3",
		"d82fcdc04c450905aa499f6980de33450f1e8d63ce83f049d37c6422be589802",
		"da40423e3190b561e295025e0733c9799637acf9019a2a69a5fa2b401cd46789",
		"df0c8d6708c3fbeadc675402ac467f84b54cace659d9810a5b5aa67a8c10311a",
		"e1b13637e63afda387771321aa48eec985f92afd5112978fbcfe97444f11796a",
		"e1f000d8f62c1d266999ab5a71c11799d7497057a7b658f98eda56d7afeb90f1",
		"e2ec4c11a6f4326eb74e65b82917f1ad82f7da44619d24ee0f0b797680e8d5c0",
		"e40a0a54ae48bb0bc18caa762f0c62f227db2190c851a6db9ac39f7b226c8a20",
		"e46ac12082a1c9403a4de56bb7e28cde86f0eb399eb97ec1fc155da09c355c98",
		"e73e181826897a995f6be659d57c38d8aed54a65f0d40d7297ba0182649695bf",
		"e93d2562ee85deab062f6f08173825e1a07984776bc1a0d3eed061cb377e0ef5",
		"e9c5659b35d03f907fbd93ac0012a3eeccc30be87b77430171d29b00b04c4149",
		"ea72ce76a0e8b5b0120bb11665ab72be6fd1c81a9aa76dc498dde509f83a7dee",
		"ede4e054659f0d4a38fa50edc718505482339315acd518f7175c8ce72db1104b",
		"f0aa75ebe9bf6549e75dfa924ea5ec79b5a5e0189b7e92f1516be41ff317d09b",
		"f5c9448d505520ffcbbd76e58eb8eeee4ed4676733bd05ba720520975e28c2d6",
		"f694a4488cef72564d5c79c4118dccf2783b6d7247e92da328b0507020095b2b",
		"f886d8b2ef0213745a426bed1e2927eadc8eee12f0c68efd227783f7e2f7004b",
		"fc365ed3f54fb8ffb717fd25963839c9e1d8c7dc480b64654fe893bc27f9e54c",
		"fd1a6796d45ab31a169e27077a8d40faad046e059d0410b33dfbbfb1e6c847dc",
		"ff708502fa6498ad38e759dc7271f6599e9f381030d6805b803c9096625a2e8e",
	}
)

type SongData struct {
	Md5        string
	Sha256     string
	Title      string
	SubTitle   string
	Genre      string
	Artist     string
	SubArtist  string
	Tag        string
	Path       string
	Folder     string
	StageFile  string
	Banner     string
	BackBmp    string
	Preview    string
	Parent     string
	Level      int32
	Difficulty int32
	MaxBpm     int32
	MinBpm     int32
	Length     int32
	Mode       int32
	Judge      int32
	Feature    int32
	Content    int32
	Date       int64
	Favorite   int32
	AddDate    int64
	Notes      int32
	ChartHash  string
}

type Folder struct {
	Title    string
	SubTitle null.String `db:"subtitle"`
	Command  null.String
	Path     string
	Banner   null.String
	Parent   string
	Type     int
	Date     int
}

func main() {
	var pdata string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewFilePicker().
				Description("Your songdata.db path:").
				Value(&pdata),
		),
	)

	if err := form.Run(); err != nil {
		panic(err)
	}

	log.Infof("songdata path=%s", pdata)

	// Load everything
	songs, err := DirectlyLoadTable[SongData](pdata, "song")
	if err != nil {
		panic(err)
	}

	illegalTitles := make([]string, 0)
	for _, v := range songs {
		for _, hash := range illegal_hash {
			if v.Md5 == hash || v.Sha256 == hash {
				illegalTitles = append(illegalTitles, v.Title)
			}
		}
	}

	if len(illegalTitles) == 0 {
		log.Info("No illegal bms found, exit")
		return
	}

	folders, err := DirectlyLoadTable[Folder](pdata, "folder")
	if err != nil {
		panic(err)
	}

	for _, v := range folders {
		for _, title := range illegalTitles {
			if v.Title == title {
				log.Warnf("Suspicious song [%s] located at [%s]", title, v.Path)
			}
		}
	}
}

func DirectlyLoadTable[T interface{}](path string, tb string) ([]*T, error) {
	db := sqlx.MustOpen("sqlite3", path)
	defer db.Close()

	rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM %s", tb))
	if err != nil {
		return nil, err
	}

	ret := make([]*T, 0)
	for rows.Next() {
		var obj T
		err = rows.StructScan(&obj)
		if err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}
	return ret, nil
}
