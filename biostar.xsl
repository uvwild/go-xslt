<?xml version='1.0'  encoding="ISO-8859-1"?>
<xsl:stylesheet
        xmlns:xsl='http://www.w3.org/1999/XSL/Transform'
        version='1.0'
>

    <xsl:output method='text' encoding="UTF-8"/>


    <xsl:template match="/">
        <xsl:apply-templates select="//PubmedArticle"/>
    </xsl:template>

    <xsl:template match="PubmedArticle" >
        <xsl:text>TY  - JOUR
        </xsl:text>
        <xsl:apply-templates select="MedlineCitation/Article/AuthorList/Author"/>
        <xsl:apply-templates select="MedlineCitation/Article/ArticleTitle"/>
        <xsl:text>
        </xsl:text>
    </xsl:template>

    <xsl:template match="ArticleTitle" >
        <xsl:text>TI  - </xsl:text>
        <xsl:value-of select="."/>
        <xsl:text>
        </xsl:text>
    </xsl:template>

    <xsl:template match="Author" >
        <xsl:text>AU  - </xsl:text>
        <xsl:value-of select="LastName"/><xsl:text> </xsl:text><xsl:value-of select="Initials"/>
        <xsl:text>
        </xsl:text>
    </xsl:template>

</xsl:stylesheet>