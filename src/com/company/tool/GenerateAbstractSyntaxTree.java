package com.company.tool;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Arrays;
import java.util.List;

public class GenerateAbstractSyntaxTree {
    public static void main(String[] args) throws IOException{
        if (args.length != 1){
            System.err.println("Usage: generate_gst <output directory>");
            System.exit(64);
        }

        String outputDir = args[0];
        defineAbstractSyntaxTree(outputDir, "Expression", Arrays.asList(
                "Binary   : Expression left, Token operator, Expression right",
                "Grouping : Expression expression",
                "Literal  : Object value",
                "Unary    : Token operator, Expression right"
        ));
    }

    private static void defineAbstractSyntaxTree(
            String outputDir, String baseName, List<String> types)
            throws IOException {
        String path = outputDir + "/" + baseName + ".java";
        PrintWriter writer = new PrintWriter(path, "UTF-8");

        writer.println("package com.company;");
        writer.println();
        writer.println("import java.util.List;");
        writer.println("import com.company.Token;");
        writer.println();
        writer.println("abstract class " + baseName + " {");
        writer.println();

        //define visitor for visitor pattern
        // Generates the visitor interface  ...for the visitor pattern
        defineVisitor(writer, baseName, types);

        //add some comments for explanation
        writer.println("/**\n" +
                " * Representing code\n" +
                " *\n" +
                " * After SCANNER has sectioned our code into tokens, we have to know how to make sense of what has been\n" +
                " * given, there in lies the SYNTACTIC GRAMMAR as the SCANNER has has done the LEXICAL GRAMMAR\n" +
                " * for us to create SYNTACTIC GRAMMAR the below is derived below and dwe create classes and subclasses from\n" +
                " * the derived formular below\n" +
                " *\n" +
                " *\n" +
                " *\n" +
                " * expression     → literal\n" +
                " *                | unary\n" +
                " *                | binary\n" +
                " *                | grouping ;\n" +
                " *\n" +
                " * literal        → NUMBER | STRING | \"true\" | \"false\" | \"nil\" ;\n" +
                " * grouping       → \"(\" expression \")\" ;\n" +
                " * unary          → ( \"-\" | \"!\" ) expression ;\n" +
                " * binary         → expression operator expression ;\n" +
                " * operator       → \"==\" | \"!=\" | \"<\" | \"<=\" | \">\" | \">=\"\n" +
                " *                | \"+\"  | \"-\"  | \"*\" | \"/\" ;\n" +
                " *\n" +
                " */");

        for(String type : types){
            String className = type.split(":")[0].trim();
            String fields = type.split(":")[1].trim();
            // Generate code for each expression type
            defineType(writer, baseName, className, fields);
            writer.println();
        }

        writer.println();
        writer.println("  abstract <R> R accept(Visitor<R> visitor);");

        writer.println("}");
        writer.close();

    }

    private static void defineVisitor(
            PrintWriter writer, String baseName, List<String> types) {
        writer.println("  interface Visitor<R> {");

        for(String type : types){
            String typeName = type.split(":")[0].trim();
            writer.println("    R visit" + typeName + baseName + "(" + typeName + " " +
                    baseName.toLowerCase() + ");");
        }

        writer.println("  }");

    }

    private static void defineType(
            PrintWriter writer, String baseName, String className, String fieldList) {
        writer.println("  static class " + className + " extends " + baseName + " {");

        //Constructor of the Expression
        writer.println("    " + className + "(" + fieldList + ") {");

        String[] fields = fieldList.split(", ");
        // Generate code for each field for the expression type
        for(String field : fields){
            String nameOfField = field.split(" ")[1];
            writer.println("      this." + nameOfField + " = " + nameOfField + ";");
        }

        writer.println("    }");

        writer.println();
        for (String field : fields) {
            writer.println("    final " + field + ";");
        }

        // add the interface implementation here
        writer.println();
        writer.println("    @Override");
        writer.println("    <R> R accept(Visitor<R> visitor) {");
        writer.println("      return visitor.visit" +
                className + baseName + "(this);");
        writer.println("    }");

        writer.println("  }");
    }
}
