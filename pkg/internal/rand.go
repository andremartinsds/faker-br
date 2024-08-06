package internal

import (
	"math/rand"
	"strconv"
)

func CepRange() string {
	cep := []string{
		"69900000", "91514420",
		"57000000", "68900000", "69000000", "40000000", "60000000",
		"70000000", "29000000", "74000000", "65000000", "78000000",
		"79000000", "30000000", "66000000", "58000000", "80000000",
		"69000000", "74000000", "74000000", "29000000", "29000000",
		"66000000", "58000000", "80000000", "69000000", "40000000",
		"60000000", "57000000", "69000000", "65000000", "30000000",
		"66000000", "58000000", "80000000", "71748506", "90718634",
		"49358280", "60536817", "43606963", "80350101", "26848205",
		"20204358", "27919210", "25615202", "48856665", "46226739",
		"90761822", "36108508", "96398708", "64705706", "90693146",
		"99895371", "49117773", "89161953", "24537786", "68533285",
		"25627359", "96785120", "70531153", "43679793", "19996835",
		"27462197", "70532572", "40611535", "64584139", "79602825",
		"29324908", "78567902", "28825523", "68143351", "38741317",
		"34835634", "68345932", "49629635", "33899878", "37444524",
		"81452974", "52817184", "83823579", "42809493", "94204702",
		"87448949", "36895724", "35573528", "95605783", "48573139",
		"78470853", "67980625", "68239213", "26668804", "38306684",
		"54597751", "88421129", "42672612", "64589536", "38848981",
		"93556834", "33509380", "48792841", "37847642", "19783218",
		"42873523", "35219486", "92352845", "83609709", "27894921",
		"78592520", "37153931", "37288289", "65499940", "53212105",
		"17432834", "94743486", "33790749", "35956857", "81326832",
		"93237408", "87917801", "63656841", "34660810", "69697930",
		"24221195", "53507539", "91410809", "26994395", "57974629",
		"72472301", "78800459", "71549649", "49564358", "59395687",
		"98434910", "97398192", "47867637", "82567719", "87502362",
		"93867173", "68563750", "91582398", "63383253", "76595173",
		"81473497", "42632643", "88568509", "27788598", "62417159",
		"31564248", "98678576", "59490593", "37536472", "32800796",
		"48637974", "34795829", "69887692", "48399739", "21848764",
		"62404792", "64668571", "57201486", "39289168", "69832819",
		"94227831", "48564402", "19454735", "16548794", "14695659",
		"72818271", "95498401", "83490582", "16453158", "71804348",
		"53963521", "92621397", "14368895", "35119194", "71354610",
		"94585194", "31667529", "35606194", "62860873", "52490651",
		"28451197", "52625487", "16767903", "96488719", "26738206",
		"17890781", "98754340", "57229593", "14315872", "35947301",
		"58139734", "47867143", "87672498", "56119594", "19613608",
		"18475763", "41785618", "42772290", "39469487", "17609837",
		"45342792", "96457201", "71205276", "56326732", "18975735",
		"81379195", "27674148", "69217798", "38718309", "26352307",
		"94336185", "39160781", "38604591", "85316265", "15832613",
		"51614578", "58792207", "89690518", "15284673", "75290973",
		"96485579", "13436314", "76989150", "68764420", "53671739",
		"34861174", "83629908", "84147360", "61709476", "67931298",
		"25484215", "79157936", "13917173", "79425174", "21795645",
		"19813573", "54657647", "97256325", "18479596", "28731392",
		"38941840", "39254130", "54286527", "81492508", "27463895",
		"42527871", "54792315", "43140402", "14982874", "46954921",
		"56984640", "92419135", "49693702", "96862175", "96413183",
		"94359270", "23469298", "47539108", "67856124", "67359625",
		"39280631", "56835612", "17948820", "28568298", "96875591",
		"24714186", "74921760", "35946735", "19481180", "46593841",
		"19320584", "86720314", "76347276", "25914489", "51483364",
		"75603305", "89290289", "65459637", "84370510", "96510109",
		"21873698", "82367135", "31219590", "51826476", "46374917",
		"24853270", "61492263", "26836908", "35216490", "82314108",
		"27815273", "36891574", "68326840", "79867345", "57329143",
		"63714201", "84985420", "48149643", "57365235", "93297124",
		"95793372", "47639617", "95314320", "13705982", "26847631",
		"78231516", "93205709", "65719573", "71643752", "67158924",
		"64739847", "39602135", "26523302", "89517708", "48526607",
		"73901827", "72537689", "74283476", "83650284", "54832498",
		"59406128", "97514706", "57624260", "78119395", "91217469",
		"18562324", "43852836", "14532697", "89751683", "27435643",
		"76197325", "18956107", "32896734", "54214680", "26842196",
		"19504520", "42613709", "83628753", "27984593", "91704306",
		"67914857", "57426173", "27415620", "91376723", "75932730",
		"32854310", "73264982", "73890169", "38205175", "54823501",
		"61827352", "86725314", "94256809", "94257369", "95386704",
		"53974589", "58309293", "63259390", "24870193", "59302398",
		"51964290", "37468561", "25489347", "15609785", "45397648",
		"86253725", "97290182", "63105576", "74283503", "43120395",
		"27195314", "84307578", "16814685", "75601802", "18407701",
		"62402139", "39625743", "19348782", "65230904", "24851726",
		"74960896", "19827278", "96469730", "17680295", "56758479",
		"64326584", "37876709", "57419920", "94610193", "19201526",
		"57627294", "37614526", "64285629", "24987139", "43921129",
		"49843193", "43941901", "57918503", "18619301", "43596284",
		"36713513", "59728860", "24964741", "93426870", "25491764",
		"53970798", "19839628", "28725870", "53827410", "29430260",
		"78230512", "26919401", "74563572", "34956823", "47358621",
		"19436987", "36238618", "86402390", "74208478", "34716169",
		"45961837", "47893506", "29816186", "48726971", "27904178",
		"67218720", "24762402", "46531386", "43179508", "18246107",
		"24603214", "76184196", "51940675", "16450609", "39417758",
		"45379512", "54703164", "94165832", "35839128", "36957142",
		"27139102", "19235960", "56315590", "53901409", "17508536",
		"49718306", "29159536", "74636629", "45136478", "96482471",
		"15206675", "93895703", "93529765", "57826702", "18973531",
		"74914675", "84308138", "34518784", "29780289", "19673102",
		"52976163", "14769178", "83216314", "47926620", "26927574",
		"64209413", "31912809", "47862356", "17358302", "53107126",
		"92736307", "83561296", "29460783", "21849403", "14856172",
		"95387673", "71913302", "28945793", "59708301", "92817713",
		"82738921", "96294819", "93489376", "16738572", "51768480",
		"18509786", "67189379", "25874670", "94205379", "15729593",
		"19860598", "25904986", "82492736", "75320180", "73829301",
		"41982148", "97587870", "73971639", "39751568", "76418471",
		"51760791", "73805472", "28425479", "46978620", "42789208",
		"94563215", "38945286", "59436194", "71318801", "69209176",
		"27968478", "23470264", "46507760", "54312869", "13924304",
		"57801143", "29749329", "29710239", "67597892", "59376235",
		"56921620", "29670409", "75847193", "13957786", "71215480",
		"31502397", "96136158", "46935496", "34768194", "62835839",
		"57643918", "47968396", "15946420", "83591489", "53407731",
		"69357704", "94313714", "47132315", "98290269", "32872708",
		"57914682", "74510329", "29631193", "17648392", "27315493",
		"32894790", "21798412", "73512908", "74910394", "53816436",
		"84675728", "31283359", "43810795", "32472591", "38645193",
	}

	return cep[rand.Intn(len(cep))]
}
func CidadeRange() string {
	cidade := []string{
		"São Paulo", "Rio de Janeiro", "Belo Horizonte", "Porto Alegre", "Curitiba", "Salvador",
		"Recife", "Fortaleza", "Brasília", "Manaus", "Belém", "Goiânia", "Campinas", "São Luís",
		"Maceió", "Natal", "Teresina", "João Pessoa", "Aracaju", "Cuiabá", "Campo Grande",
		"Vitória", "Florianópolis", "Macapá", "Palmas", "Boa Vista", "Rio Branco", "Santos",
		"Sorocaba", "Ribeirão Preto", "Uberlândia", "São José dos Campos", "Mogi das Cruzes",
		"Piracicaba", "Jundiaí", "Bauru", "Londrina", "Maringá", "Cascavel", "Foz do Iguaçu",
		"Joinville", "Blumenau", "Chapecó", "Caxias do Sul", "Pelotas", "Santa Maria", "Canoas",
		"Gravataí", "Novo Hamburgo", "São Leopoldo", "Niterói", "Duque de Caxias", "São Gonçalo",
		"Volta Redonda", "Petrópolis", "Campos dos Goytacazes", "Vila Velha", "Cariacica",
		"Serra", "Cachoeiro de Itapemirim", "Montes Claros", "Uberaba", "Divinópolis", "Ipatinga",
		"Sete Lagoas", "Betim", "Contagem", "Anápolis", "Aparecida de Goiânia", "Rondonópolis",
		"Sinop", "Dourados", "Três Lagoas", "Corumbá", "Mossoró", "Caucaia", "Maracanaú",
		"Sobral", "Juazeiro do Norte", "Petrolina", "Cabo de Santo Agostinho", "Olinda",
		"Caruaru", "Jaboatão dos Guararapes", "Paulista", "Lauro de Freitas", "Feira de Santana",
		"Vitória da Conquista", "Camaçari", "Barreiras", "Itabuna", "Ilhéus", "Jequié", "Teixeira de Freitas",
		"Porto Seguro", "Santarém", "Marabá", "Castanhal", "Parauapebas", "Ananindeua", "Abaetetuba",
		"Tucuruí", "Altamira", "Araguaína", "Palmas", "Gurupi", "Imperatriz", "São José de Ribamar",
		"Timon", "Caxias", "Bacabal", "Parnaíba", "Picos", "Piripiri", "Floriano", "Crato", "Arapiraca",
		"Palmeira dos Índios", "Abaetetuba", "Bacabal", "Codó", "Santa Inês", "Caxias", "Timon", "Parnaíba",
		"Floriano", "Piripiri", "Picos", "Luziânia", "Águas Lindas de Goiás", "Formosa", "Planaltina",
		"Santo Antônio do Descoberto", "Novo Gama", "Valparaíso de Goiás", "Cidade Ocidental", "Cristalina",
		"Catalão", "Itumbiara", "Jataí", "Rio Verde", "Gurupi", "Porto Nacional", "Araguaína", "Colinas do Tocantins",
		"Paraíso do Tocantins", "Guaraí", "Miracema do Tocantins", "Tocantinópolis", "Araguatins", "Xambioá",
		"Arraias", "Goiatins", "Barra do Garças", "Cáceres", "Rondonópolis", "Tangará da Serra", "Sinop", "Alta Floresta",
		"Lucas do Rio Verde", "Sorriso", "Primavera do Leste", "Campo Verde", "Jaciara", "Nobres", "Rosário Oeste",
		"Chapada dos Guimarães", "Poconé", "Barra do Bugres", "Pontes e Lacerda", "Juína", "Colniza", "Vilhena", "Ariquemes",
		"Ji-Paraná", "Rolim de Moura", "Guajará-Mirim", "Cacoal", "Pimenta Bueno", "Presidente Médici", "Machadinho d'Oeste",
		"Ouro Preto do Oeste", "Buritis", "Alto Paraíso", "Alvorada do Oeste", "Espigão d'Oeste", "Alta Floresta d'Oeste",
		"Nova Mamoré", "Campo Novo de Rondônia", "Mirante da Serra", "Monte Negro", "Santa Luzia d'Oeste", "Seringueiras",
		"Theobroma", "Vale do Anari", "Vale do Paraíso", "Vila Nova do Mamoré", "Costa Marques", "Cujubim", "Governador Jorge Teixeira",
		"Jaru", "São Miguel do Guaporé", "São Francisco do Guaporé", "Itapuã do Oeste", "Nova União", "Cacaulândia", "Chupinguaia",
		"Nova Brasilândia d'Oeste", "Rio Crespo", "São Felipe d'Oeste", "São Francisco do Guaporé", "São Miguel do Guaporé",
		"Cerejeiras", "Corumbiara", "Pimenteiras do Oeste", "Cabixi", "Colorado do Oeste", "São Domingos do Guaporé", "São João da Baliza",
		"São João do Sul", "São Luiz", "São Miguel do Tocantins", "São Paulo de Olivença", "São Pedro da Água Branca", "São Raimundo Nonato",
		"São Sebastião do Uatumã", "São Sebastião do Caí", "São Sebastião do Maranhão", "São Sebastião do Oeste", "São Sebastião do Paraíso",
		"São Sebastião do Passé", "São Sebastião do Rio Verde", "São Sebastião do Tocantins", "São Sepé", "São Simão", "São Tomé",
		"São Tomé das Letras", "São Tomás de Aquino", "São Valentim", "São Valentim do Sul", "São Valério", "São Valério da Natividade",
		"São Vendelino", "São Vicente Ferrer", "São Vicente do Sul", "São Vicente de Minas", "São Vicente de Paula", "São Vicente de Carvalho",
		"São Vicente de Cima", "São Vicente de Fora", "São Vicente de Pouso Alegre", "São Vicente de Xingu", "São Vicente do Araguaia",
		"São Vicente do Ivaí", "São Vicente do Paul", "São Vicente do Prata", "São Vicente do Ribeira", "São Vicente do Sul", "São Vicente do Turvo",
		"São Vicente do Xingu", "São Vicente do Pontal", "São Vicente de Minas", "São Vicente de Paula", "São Vicente de Fora", "São Vicente de Pouso Alegre",
		"São Vicente de Xingu", "São Vicente do Araguaia", "São Vicente do Ivaí", "São Vicente do Paul", "São Vicente do Prata", "São Vicente do Ribeira",
		"São Vicente do Sul", "São Vicente do Turvo", "São Vicente do Xingu", "São Vicente do Pontal", "São Vicente de Minas", "São Vicente de Paula",
		"São Vicente de Fora", "São Vicente de Pouso Alegre", "São Vicente de Xingu", "São Vicente do Araguaia", "São Vicente do Ivaí", "São Vicente do Paul",
		"São Vicente do Prata", "São Vicente do Ribeira", "São Vicente do Sul", "São Vicente do Turvo", "São Vicente do Xingu", "São Vicente do Pontal",
		"São Vicente de Minas", "São Vicente de Paula", "São Vicente de Fora", "São Vicente de Pouso Alegre", "São Vicente de Xingu", "São Vicente do Araguaia",
		"São Vicente do Ivaí", "São Vicente do Paul", "São Vicente do Prata", "São Vicente do Ribeira", "São Vicente do Sul", "São Vicente do Turvo",
		"São Vicente do Xingu", "São Vicente do Pontal", "São Vicente de Minas", "São Vicente de Paula", "São Vicente de Fora", "São Vicente de Pouso Alegre",
		"São Vicente de Xingu", "São Vicente do Araguaia", "São Vicente do Ivaí", "São Vicente do Paul", "São Vicente do Prata", "São Vicente do Ribeira",
		"São Vicente do Sul", "São Vicente do Turvo", "São Vicente do Xingu", "São Vicente do Pontal", "São Vicente de Minas", "São Vicente de Paula",
		"São Vicente de Fora", "São Vicente de Pouso Alegre", "São Vicente de Xingu", "São Vicente do Araguaia", "São Vicente do Ivaí", "São Vicente do Paul",
		"São Vicente do Prata", "São Vicente do Ribeira",
	}

	return cidade[rand.Intn(len(cidade))]
}
func EstadoRange() string {
	estado := []string{
		"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Distrito Federal",
		"Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
		"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí",
		"Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia",
		"Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins",
	}
	return estado[rand.Intn(len(estado))]
}
func BairroRange() string {
	bairro := []string{
		"Centro", "Jardim Paulista", "Copacabana", "Ipanema", "Leblon", "Botafogo",
		"Barra da Tijuca", "Tijuca", "Flamengo", "Glória", "Lapa", "Santa Teresa",
		"Moema", "Pinheiros", "Vila Madalena", "Vila Mariana", "Santana", "Tatuapé",
		"Bela Vista", "Lapa", "Perdizes", "Consolação", "Brooklin", "Morumbi",
		"Jardim Europa", "Higienópolis", "Paraíso", "Cambuci", "Aclimação",
		"Vila Clementino", "Jardins", "Brás", "Bom Retiro", "Liberdade", "Pacaembu",
		"Itaim Bibi", "Vila Olímpia", "Saúde", "Santo Amaro", "Itaquera", "Vila Prudente",
		"São Mateus", "Bangu", "Realengo", "Campo Grande", "Pavuna", "Madureira",
		"São Cristóvão", "Bento Ribeiro", "Guaratiba", "Rocha Miranda", "Pilares",
		"Jacarepaguá", "Meier", "Engenho Novo", "Ilha do Governador", "Santa Cruz",
		"Cosme Velho", "Laranjeiras", "Gávea", "Vidigal", "Vila Isabel", "Estácio",
		"São Conrado", "Urca", "Recreio dos Bandeirantes", "Sepetiba", "Vargem Grande",
		"Vargem Pequena", "Maracanã", "Andaraí", "Grajaú", "Barra Funda", "Água Branca",
		"Jardim América", "Belém", "Belem Novo", "Moinhos de Vento", "Navegantes", "Menino Deus",
		"Petrópolis", "Teresópolis", "Santo Antônio", "Auxiliadora", "Tristeza", "Vila Assunção",
		"Cristal", "Ipanema", "Vila Nova", "Restinga", "Vila Ipiranga", "São João", "Sarandi",
		"Passo da Areia", "Boa Vista", "Santa Cecília", "São Geraldo", "Glória",
	}
	return bairro[rand.Intn(len(bairro))]
}
func NumeroPorTamanho(size int) int {
	chars := "0123456789"
	numero := make([]byte, size)
	for i, _ := range numero {
		numero[i] = chars[rand.Intn(len(chars))]
	}
	n, _ := strconv.Atoi(string(numero))
	return n
}
func ComplementoRange() string {
	complemento := []string{
		"Apto 101", "Apto 202", "Apto 303", "Apto 404", "Apto 505", "Apto 606",
		"Bloco A", "Bloco B", "Bloco C", "Bloco D", "Bloco E", "Bloco F",
		"Fundos", "Casa 1", "Casa 2", "Casa 3", "Casa 4", "Casa 5",
		"Casa A", "Casa B", "Casa C", "Casa D", "Casa E", "Casa F",
		"Loja 1", "Loja 2", "Loja 3", "Loja 4", "Loja 5", "Loja 6",
		"Loja A", "Loja B", "Loja C", "Loja D", "Loja E", "Loja F",
		"Sala 101", "Sala 102", "Sala 201", "Sala 202", "Sala 301", "Sala 302",
		"Cobertura", "Térreo", "Sobreloja", "Galpão 1", "Galpão 2", "Galpão 3",
		"Anexo 1", "Anexo 2", "Anexo 3", "Anexo 4", "Anexo 5", "Anexo 6",
		"Kitnet 101", "Kitnet 102", "Kitnet 103", "Kitnet 104", "Kitnet 105", "Kitnet 106",
	}
	return complemento[rand.Intn(len(complemento))]
}
func ObservacaoRange() string {
	obs := []string{
		"Portão azul", "Próximo ao supermercado", "Entrada pela lateral",
		"Ao lado da padaria", "Frente para a praça", "Próximo ao ponto de ônibus",
		"Acesso pelo portão 2", "Entrada pela garagem", "Prédio amarelo",
		"Casa com muro alto", "Ao lado da farmácia", "Próximo à escola",
		"Em frente ao parque", "Portão verde", "Ao lado da igreja",
		"Próximo ao hospital", "Casa de esquina", "Portão branco",
		"Acesso pela rua dos fundos", "Casa com jardim na frente",
		"Próximo ao banco", "Em frente à delegacia", "Ao lado do shopping",
		"Casa com varanda", "Próximo ao clube", "Frente para a avenida principal",
		"Portão de ferro", "Próximo ao posto de gasolina", "Ao lado da academia",
		"Casa com piscina", "Próximo à estação de metrô", "Em frente ao correio",
		"Ao lado do restaurante", "Casa com garagem dupla", "Próximo ao cinema",
		"Casa com portão automático", "Em frente ao mercado", "Ao lado do parque",
		"Casa com churrasqueira", "Próximo à biblioteca", "Ao lado da universidade",
		"Casa com jardim de inverno", "Próximo ao teatro", "Em frente ao ginásio",
		"Ao lado do salão de beleza", "Próximo à pizzaria", "Casa com telhado vermelho",
		"Em frente ao hotel", "Ao lado da agência de viagens", "Próximo ao centro comunitário",
		"Casa com portão de madeira", "Em frente à lotérica",
	}
	return obs[rand.Intn(len(obs))]
}
func UfRange() string {
	u := []string{
		"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG",
		"PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO",
	}
	return u[rand.Intn(len(u))]
}
