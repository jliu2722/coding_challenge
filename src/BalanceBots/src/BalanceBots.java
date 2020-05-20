import java.io.*;
import java.net.URL;
import java.util.*;
import static org.junit.Assert.*;

public class BalanceBots {
    static Map<String, List<String>> botMap = new HashMap<>();
    static Map<String, List<String>> outputMap = new HashMap<>();
    static Map<String, Distribution> ruleMap = new HashMap<>();

    static class Distribution{
        String low;
        String high;
        Distribution(String low, String high){
            this.low = low;
            this.high = high;
        }
    }

    static void assignChip(String line){
        String chip = line.split(" goes to ")[0];
        String bot = line.split(" goes to ")[1];
        if (!botMap.containsKey(bot)) {
            botMap.put(bot, List.of(chip));
        } else{
            List<String> chips = new ArrayList<String>(botMap.get(bot));
            chips.add(chip);
            botMap.put(bot, chips);
        }
    }

    static void addBotIfNotThere(String botOrOutut){
        if (botOrOutut.startsWith("bot") && !botMap.containsKey(botOrOutut)) {
            botMap.put(botOrOutut, new ArrayList<>());
        }
    }

    static void assignRule(String line){
        String botId = line.split(" gives ")[0];
        String low = line.split(" low to ")[1].split(" and high ")[0];
        String high = line.split(" high to ")[1];
        addBotIfNotThere(low);
        addBotIfNotThere(high);
        ruleMap.put(botId, new Distribution(low, high));
    }

    static void parseAct(String line) {
        if (line.startsWith("value")) {
            assignChip(line);
        } else {
            assignRule(line);
        }
    }

    static void moveChip(String chip, String source, String destination){
        List<String> list = botMap.get(source);
        list.remove(chip);
        botMap.put(source, list);
        if (destination.startsWith("bot")) {
            list = new ArrayList<String>(botMap.get(destination));
            list.add(chip);
            botMap.put(destination, list);
        } else if (destination.startsWith("output")){
            if (outputMap.containsKey(destination)){
                list = new ArrayList<String>(outputMap.get(destination));
                list.add(chip);
                botMap.put(source, list);
            } else {
                outputMap.put(destination, List.of(chip));
            }
        }
    }

    static void runRules(){
        for (String bot: botMap.keySet()){
            if (botMap.get(bot).size() > 1) {
                String chip0 = botMap.get(bot).get(0);
                String chip1 = botMap.get(bot).get(1);
                if (chip0.compareTo(chip1) < 0) {
                    moveChip(chip0, bot, ruleMap.get(bot).low);
                    moveChip(chip1, bot, ruleMap.get(bot).high);
                } else {
                    moveChip(chip1, bot, ruleMap.get(bot).low);
                    moveChip(chip0, bot, ruleMap.get(bot).high);
                }
            }
        }
    }

    public static void main(String[] args) throws IOException {
        URL path = ClassLoader.getSystemResource("inst.txt");
        File f = new File(path.getFile());
        BufferedReader reader = new BufferedReader(new FileReader(f));
        String line = null;
        while ((line = reader.readLine()) != null) {
            System.out.println(line);
            parseAct(line);
        }
        runRules();

        //unit tests below to valdiate the answers
        assertEquals(0, botMap.get("bot 2").size());
        assertEquals(0, botMap.get("bot 1").size());
        assertEquals(0, botMap.get("bot 0").size());
        assertEquals("value 5", outputMap.get("output 0").get(0));
        assertEquals("value 2", outputMap.get("output 1").get(0));
        assertEquals("value 3", outputMap.get("output 2").get(0));
    }
}



