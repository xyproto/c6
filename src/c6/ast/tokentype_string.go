// generated by stringer -type=TokenType token.go; DO NOT EDIT

package ast

import "fmt"

const _TokenType_name = "T_SPACET_COMMENT_LINET_COMMENT_BLOCKT_SEMICOLONT_COMMAT_IDENTT_URLT_MEDIAT_TRUET_FALSET_NULLT_ONLYT_MS_PARAM_NAMET_FUNCTION_NAMET_ID_SELECTORT_CLASS_SELECTORT_TYPE_SELECTORT_UNIVERSAL_SELECTORT_PARENT_SELECTORT_PSEUDO_SELECTORT_FUNCTIONAL_PSEUDOT_INTERPOLATION_SELECTORT_LITERAL_CONCATT_CONCATT_MS_PROGIDT_AND_SELECTORT_DESCENDANT_COMBINATORT_CHILD_COMBINATORT_ADJACENT_SIBLING_COMBINATORT_GENERAL_SIBLING_COMBINATORT_UNICODE_RANGET_IFT_ELSET_ELSE_IFT_INCLUDET_MIXINT_FUNCTIONT_RETURNT_GLOBALT_DEFAULTT_IMPORTANTT_OPTIONALT_FONT_FACET_LOGICAL_NOTT_LOGICAL_ORT_LOGICAL_ANDT_LOGICAL_XORT_NOPT_PLUST_DIVT_MULT_MINUST_MODT_BRACE_STARTT_BRACE_ENDT_LANG_CODET_BRACKET_LEFTT_ATTRIBUTE_NAMET_BRACKET_RIGHTT_EQUALT_UNEQUALT_GTT_LTT_GET_LET_ASSIGNT_ATTR_EQUALT_ATTR_TILDE_EQUALT_ATTR_HYPHEN_EQUALT_VARIABLET_IMPORTT_AT_RULET_CHARSETT_QQ_STRINGT_Q_STRINGT_UNQUOTE_STRINGT_PAREN_STARTT_PAREN_ENDT_CONSTANTT_INTEGERT_FLOATT_UNIT_NONET_UNIT_PERCENTT_UNIT_SECONDT_UNIT_MILLISECONDT_UNIT_EMT_UNIT_EXT_UNIT_CHT_UNIT_REMT_UNIT_CMT_UNIT_INT_UNIT_MMT_UNIT_PCT_UNIT_PTT_UNIT_PXT_UNIT_VHT_UNIT_VWT_UNIT_VMINT_UNIT_VMAXT_UNIT_HZT_UNIT_KHZT_UNIT_DPIT_UNIT_DPCMT_UNIT_DPPXT_UNIT_DEGT_UNIT_GRADT_UNIT_RADT_UNIT_TURNT_PROPERTY_NAME_TOKENT_PROPERTY_VALUET_HEX_COLORT_COLONT_INTERPOLATION_STARTT_INTERPOLATION_INNERT_INTERPOLATION_END"

var _TokenType_index = [...]uint16{0, 7, 21, 36, 47, 54, 61, 66, 73, 79, 86, 92, 98, 113, 128, 141, 157, 172, 192, 209, 226, 245, 269, 285, 293, 304, 318, 341, 359, 388, 416, 431, 435, 441, 450, 459, 466, 476, 484, 492, 501, 512, 522, 533, 546, 558, 571, 584, 589, 595, 600, 605, 612, 617, 630, 641, 652, 666, 682, 697, 704, 713, 717, 721, 725, 729, 737, 749, 767, 786, 796, 804, 813, 822, 833, 843, 859, 872, 883, 893, 902, 909, 920, 934, 947, 965, 974, 983, 992, 1002, 1011, 1020, 1029, 1038, 1047, 1056, 1065, 1074, 1085, 1096, 1105, 1115, 1125, 1136, 1147, 1157, 1168, 1178, 1189, 1210, 1226, 1237, 1244, 1265, 1286, 1305}

func (i TokenType) String() string {
	if i < 0 || i+1 >= TokenType(len(_TokenType_index)) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
