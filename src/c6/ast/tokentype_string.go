// generated by stringer -type=TokenType token.go; DO NOT EDIT

package ast

import "fmt"

const _TokenType_name = "T_SPACET_COMMENT_LINET_COMMENT_BLOCKT_SEMICOLONT_COMMAT_IDENTT_URLT_MEDIAT_TRUET_FALSET_NULLT_ONLYT_MS_PARAM_NAMET_FUNCTION_NAMET_ID_SELECTORT_CLASS_SELECTORT_TYPE_SELECTORT_UNIVERSAL_SELECTORT_PARENT_SELECTORT_PSEUDO_SELECTORT_FUNCTIONAL_PSEUDOT_INTERPOLATION_SELECTORT_LITERAL_CONCATT_CONCATT_MS_PROGIDT_AND_SELECTORT_DESCENDANT_COMBINATORT_CHILD_COMBINATORT_ADJACENT_SIBLING_COMBINATORT_GENERAL_SIBLING_COMBINATORT_UNICODE_RANGET_IFT_ELSET_ELSE_IFT_INCLUDET_EACHT_WHENT_MIXINT_FUNCTIONT_FORT_FOR_FROMT_FOR_THROUGHT_FOR_TOT_FOR_INT_WHILET_RETURNT_RANGET_CONTENTT_GLOBALT_DEFAULTT_IMPORTANTT_OPTIONALT_FONT_FACET_LOGICAL_NOTT_LOGICAL_ORT_LOGICAL_ANDT_LOGICAL_XORT_NOPT_PLUST_DIVT_MULT_MINUST_MODT_BRACE_STARTT_BRACE_ENDT_LANG_CODET_BRACKET_LEFTT_ATTRIBUTE_NAMET_BRACKET_RIGHTT_EQUALT_UNEQUALT_GTT_LTT_GET_LET_ASSIGNT_ATTR_EQUALT_ATTR_HYPHEN_EQUALT_INCLUDE_MATCHT_PREFIX_MATCHT_DASH_MATCHT_SUFFIX_MATCHT_SUBSTRING_MATCHT_VARIABLET_IMPORTT_AT_RULET_CHARSETT_QQ_STRINGT_Q_STRINGT_UNQUOTE_STRINGT_PAREN_STARTT_PAREN_ENDT_CONSTANTT_INTEGERT_FLOATT_UNIT_NONET_UNIT_PERCENTT_UNIT_SECONDT_UNIT_MILLISECONDT_UNIT_EMT_UNIT_EXT_UNIT_CHT_UNIT_REMT_UNIT_CMT_UNIT_INT_UNIT_MMT_UNIT_PCT_UNIT_PTT_UNIT_PXT_UNIT_VHT_UNIT_VWT_UNIT_VMINT_UNIT_VMAXT_UNIT_HZT_UNIT_KHZT_UNIT_DPIT_UNIT_DPCMT_UNIT_DPPXT_UNIT_DEGT_UNIT_GRADT_UNIT_RADT_UNIT_TURNT_PROPERTY_NAME_TOKENT_PROPERTY_VALUET_HEX_COLORT_COLONT_INTERPOLATION_STARTT_INTERPOLATION_INNERT_INTERPOLATION_END"

var _TokenType_index = [...]uint16{0, 7, 21, 36, 47, 54, 61, 66, 73, 79, 86, 92, 98, 113, 128, 141, 157, 172, 192, 209, 226, 245, 269, 285, 293, 304, 318, 341, 359, 388, 416, 431, 435, 441, 450, 459, 465, 471, 478, 488, 493, 503, 516, 524, 532, 539, 547, 554, 563, 571, 580, 591, 601, 612, 625, 637, 650, 663, 668, 674, 679, 684, 691, 696, 709, 720, 731, 745, 761, 776, 783, 792, 796, 800, 804, 808, 816, 828, 847, 862, 876, 888, 902, 919, 929, 937, 946, 955, 966, 976, 992, 1005, 1016, 1026, 1035, 1042, 1053, 1067, 1080, 1098, 1107, 1116, 1125, 1135, 1144, 1153, 1162, 1171, 1180, 1189, 1198, 1207, 1218, 1229, 1238, 1248, 1258, 1269, 1280, 1290, 1301, 1311, 1322, 1343, 1359, 1370, 1377, 1398, 1419, 1438}

func (i TokenType) String() string {
	if i < 0 || i+1 >= TokenType(len(_TokenType_index)) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
